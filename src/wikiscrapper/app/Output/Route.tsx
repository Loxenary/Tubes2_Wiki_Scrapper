import { useState } from "react";
import PathInterface from "./PathData";

const SinglePath: React.FC<PathInterface> = ({ index, item }) => (
  <div className="flex flex-row gap-5 items-center text-xl my-5">
    <div className="w-10 h-10 bg-gray-300 items-center flex justify-center rounded-full">
      {index}
    </div>
    <h1 className="">{item}</h1>
  </div>
);

const RouteOutput = () => {
  const dummyData: PathInterface[] = [
    { index: "1", item: "Example 1" },
    { index: "2", item: "Example 2" },
    { index: "3", item: "Example 3" },
  ];

  return (
    <div className="text-lg my-10">
        <div className="flex gap-5 ">
        <h1>Time Execution: </h1>
        <div className=" w-max h-7">..ms</div>
      </div>
      {dummyData.map((item, index) => (
        <SinglePath key={index} index={item.index} item={item.item} />
      ))}
    </div>
  );
};

export default RouteOutput;
