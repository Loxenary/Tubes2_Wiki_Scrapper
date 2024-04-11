"use client";
import Title from "./title";
import ToggleAlgorithm from "./toggleAlgorithm";
import EntryWiki from "./entryWiki";
import OutputPage from "../Output/page";
import { createContext, useState } from "react";
import { WikiSearchContextProvider } from "@/Context/SearchContext";
import { OutputContextProvider } from "@/Context/OutputContext";

// Interface used in BoolOutputSetup
export interface ISetupOutputPage {
  setOutputState: React.Dispatch<React.SetStateAction<boolean>>;
}

// Context that have responsibility on Visibility of the output page
export const BoolOutputSetup = createContext<ISetupOutputPage | undefined>(
  undefined
);

const InputEntry = () => {
  const [outputState, setOutputState] = useState(false);
  return (
    // This is context responsibility to hold data such as algorithm
    <WikiSearchContextProvider>
      <div className="w-full justify-center items-center flex flex-col gap-y-5">
        {/* Hold Title */}
        <Title></Title>

        {/* Hold the toggle slider */}
        <ToggleAlgorithm></ToggleAlgorithm>

        {/* Hold the Output Data from backend */}
        <OutputContextProvider>

          {/* usage : turn on or off visibility of the output page*/}
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
