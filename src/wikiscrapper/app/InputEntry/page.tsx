"use client";
import Title from "./title";
import ToggleAlgorithm from "./toggleAlgorithm";
import EntryWiki from "./entryWiki";
import OutputPage from "../Output/page";
import { createContext, useState } from "react";
import { WikiSearchContextProvider } from "@/Context/SearchContext";

export interface IOutputContext {
  setOutputState: React.Dispatch<React.SetStateAction<boolean>>;
}

export const OutputContext = createContext<IOutputContext | undefined>(undefined);

const InputEntry = () => {
  const [outputState, setOutputState] = useState(false);
  return (
    <WikiSearchContextProvider>
      <div className="w-full justify-center items-center flex flex-col gap-y-5">
        <Title></Title>
        <ToggleAlgorithm></ToggleAlgorithm>
        <OutputContext.Provider value={{ setOutputState }}>
          <EntryWiki></EntryWiki>
        </OutputContext.Provider>

        {outputState ? <OutputPage></OutputPage> : null}
      </div>
    </WikiSearchContextProvider>
  );
};

export default InputEntry;
