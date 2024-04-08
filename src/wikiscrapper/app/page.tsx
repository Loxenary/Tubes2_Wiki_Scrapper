import InputEntry from "./InputEntry/page";
import OutputPage from "./Output/page";
export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24 bg-white text-black">
      <InputEntry></InputEntry>
      <OutputPage></OutputPage>
    </main>
  );
}
