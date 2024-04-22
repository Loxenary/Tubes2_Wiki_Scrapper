import { useEffect, useState } from "react";

interface IAutoComplete {
  data: string;
  setData: (data: string) => void;
}

const Autocomplete: React.FC<IAutoComplete> = ({ data, setData }) => {
  const [results, setResults] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      if (data.trim() !== "") {
        try {
          const response = await fetch(
            `https://en.wikipedia.org/w/api.php?action=opensearch&format=json&search=${data}&origin=*`
          );
          const responseData = await response.json();
          setResults(responseData[1]);
        } catch (error) {
          console.error("Error fetching data:", error);
        }
      } else {
        setResults([]);
      }
    };

    fetchData();
  }, [data]);

  const handleSubmit = (data: string) => {
    setData(data);
  };

  return (
    <div className="flex flex-col flex-grow ">
      {results.map((result, index) => (
        <button key={index} type="button" onClick={() => handleSubmit(result)}>
          {result}
        </button>
      ))}
    </div>
  );
};

export default Autocomplete;
