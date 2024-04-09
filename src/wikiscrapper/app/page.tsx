import InputEntry from "./InputEntry/page";
import Toast from "@/components/toast";
export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center p-24 bg-white text-black">
      <InputEntry></InputEntry>
      <Toast></Toast>
    </main>
  );
}
