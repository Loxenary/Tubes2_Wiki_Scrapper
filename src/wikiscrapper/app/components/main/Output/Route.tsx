"use client";
import { useOutputContext } from "@/Context/OutputContext";
import PathInterface from "./PathData";

interface ISinglePath extends PathInterface{
  index : number;
}

const SinglePath: React.FC<ISinglePath> = ({ index, item }) => (
  <div className="flex flex-row gap-5 items-center text-xl my-5">
    <div className="w-10 h-10 bg-gray-300 items-center flex justify-center rounded-full">
      {index}
    </div>
    <h1 className="">{item}</h1>
  </div>
);

const RouteOutput = () => {
  const { time, listPath } = useOutputContext();

  const timeExecution = () => {
    return (
      <div className="flex gap-5 ">
        <h1>Time Execution: </h1>
        <div className=" w-max h-7">
          {time ? time.toString() + "" : "..ms"}
        </div>
      </div>
    );
  };

  return (
    <div className="text-lg my-10">
      {timeExecution()}
      {listPath
        ? listPath.map((item, index) => (
            <SinglePath
              key={index}
              index={index}
              item={item.item}
            ></SinglePath>
          ))
        : null}
    </div>
  );
};

export default RouteOutput;
