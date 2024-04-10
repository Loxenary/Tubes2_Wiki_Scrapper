"use client";
import Title from "./title";
import ToggleAlgorithm from "./toggleAlgorithm";
import EntryWiki from "./entryWiki";
import OutputPage from "../Output/page";
import { createContext, useState } from "react";
import { WikiSearchContextProvider } from "@/Context/SearchContext";
import { OutputContextProvider } from "@/Context/OutputContext";
export interface ISetupOutputPage {
  setOutputState: React.Dispatch<React.SetStateAction<boolean>>;
}

export const BoolOutputSetup = createContext<ISetupOutputPage | undefined>(
  undefined
);

const InputEntry = () => {
  const [outputState, setOutputState] = useState(false);
  return (
    <WikiSearchContextProvider>
      <div className="w-full justify-center items-center flex flex-col gap-y-5">
        {/* Hold Title */}
        <Title></Title>

        {/* Hold the toggle slider */}
        <ToggleAlgorithm></ToggleAlgorithm>

        {/* Hold the Output Data from backend */}
        <OutputContextProvider>

          {/* Hold the data for turn on the backend */}
          <BoolOutputSetup.Provider value={{ setOutputState }}>

            {/* Hold Component for User input */}
            <EntryWiki></EntryWiki>
          </BoolOutputSetup.Provider>

          {/* Hold the Output Components */}
          {outputState ? <OutputPage></OutputPage> : null}
        </OutputContextProvider>
      </div>
    </WikiSearchContextProvider>
  );
};

export default InputEntry;
