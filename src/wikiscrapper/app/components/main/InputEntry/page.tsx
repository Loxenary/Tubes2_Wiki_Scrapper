"use client";

import ToggleAlgorithm from "./toggleAlgorithm";
import EntryWiki from "./entryWiki";
import { createContext, useState } from "react";
import { WikiSearchContextProvider } from "@/Context/SearchContext";

// Interface used in BoolOutputSetup


const InputEntry = () => {
  
  return (
    // This is context responsibility to hold data such as algorithm
    <WikiSearchContextProvider>
      <div
      className="flex flex-col items-center justify-center pt-40"
      id="race"
      >
        {/* Hold Title */}
        <div className="max-w-[1240px] mx-10 grid md:grid-cols-2 gap-10 gap-y-5">
          <div data-aos="fade-right">
            <h1 className="text-7xl font-semibold text-transparent bg-clip-text bg-white md:py-20 text-center md:text-left md:text-9xl">
              LET'S RACE
            </h1>
          </div>

          <div className="py-10">
            {/* Hold the toggle slider */}
            <ToggleAlgorithm></ToggleAlgorithm>        
            {/* Hold Component for User input */}
            <EntryWiki></EntryWiki>
          </div>
        </div>

      </div>
    </WikiSearchContextProvider>
  );
};

export default InputEntry;
