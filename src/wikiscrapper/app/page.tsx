"use client";
import Hero from "./components/main/Hero";
import About from "./components/main/About";
import { useEffect, useState, createContext } from "react";

import AOS from "aos";
import "aos/dist/aos.css";

import InputEntry from "./components/main/InputEntry/page";
import Toast from "@/app/components/main/toast";
import OutputPage from "./components/main/Output/page";
import { OutputContextProvider } from "@/Context/OutputContext";

export interface ISetupOutputPage {
  setOutputState:  React.Dispatch<React.SetStateAction<boolean>>
}

// Context that have responsibility on Visibility of the output page
export const BoolOutputSetup = createContext<ISetupOutputPage>({
  setOutputState: () => {},
});

export default function Home() {
  useEffect(() => {
    AOS.init({duration:1200})
  })
  const [outputState, setOutputState] = useState(false);
  return (
    <main className='h-full w-full'>
      <div className='flex flex-col gap-20 z-20'>
        <Hero />
        
        <OutputContextProvider>
        <BoolOutputSetup.Provider value={{ setOutputState }}>
          <InputEntry></InputEntry>
        </BoolOutputSetup.Provider>
        
        {/* Hold the Output Data from backend */}
          {
            outputState? <OutputPage></OutputPage> : null
          }
        </OutputContextProvider>
        <Toast></Toast>
        <About />
      </div>
    </main>

    // {/* <main className="flex min-h-screen flex-col items-center p-24 bg-white text-black">
    // </main> */}
  );
}
