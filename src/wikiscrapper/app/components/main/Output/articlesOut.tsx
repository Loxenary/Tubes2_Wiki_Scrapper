"use client";
import { useOutputContext } from "@/Context/OutputContext";

const ArticlesOut = () => {
  const { checkcount, numpassed } = useOutputContext();

  const articles = (name: string, data: any) => {
    return (
      <div className="flex gap-5 justify-center items-center">
      <h1>{name}</h1>
      <div className="bg-[#6147df]/20 border border-[#6147df] shadow-lg shadow-[#6147df]/50 text-white flex rounded-md justify-center items-center px-2">
        {data ? data.toString() : ""}
      </div>
      </div>
    );
  };

  return (
    <div className="justify-between items-center text-2xl flex flex-wrap gap-y-10 gap-x-20">
      {articles("Articles Checked", checkcount)}
      {articles("Articles Passed", numpassed)}
    </div>
  );
};

export default ArticlesOut;
