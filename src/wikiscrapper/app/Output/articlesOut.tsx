"use client";
import { useOutputContext } from "@/Context/OutputContext";

const ArticlesOut = () => {
  const { checkcount, numpassed } = useOutputContext();

  const articles = (name: string, data: any) => {
    return (
      <div className="flex gap-5 justify-center items-center">
        <h1>{name}</h1>
        <div className="bg-gray-500 w-10 h-10 text-white flex justify-center items-center">
          {data ? data.toString() : ""}
        </div>
      </div>
    );
  };

  return (
    <div className="justify-between items-center text-2xl flex gap-x-20">
      {articles("Articles Checked", checkcount)}
      {articles("Articles Passed", numpassed)}
    </div>
  );
};

export default ArticlesOut;
