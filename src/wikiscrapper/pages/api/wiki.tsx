export async function WikipediaExistChecker(title: string): Promise<string[]> {
  try {
    const url = `https://en.wikipedia.org/w/api.php?action=opensearch&limit=1&profile=strict&format=json&search=${title}&origin=*`;
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(
        `Failed to fetch Wikipedia search results: ${response.statusText}`
      );
    }

    const searchData: any[] = await response.json();
    const titles: string[] = searchData[1] || [];

    return titles;
  } catch (error) {
    console.error("Error fetching Wikipedia search results:", error);
    throw error;
  }
}
