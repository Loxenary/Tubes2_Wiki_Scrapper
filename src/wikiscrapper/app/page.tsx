"use client";

import Image from "next/image";
import Head from 'next/head';
import Hero from "./components/main/Hero";
import About from "./components/main/About";
import { useEffect } from "react";

import AOS from "aos";
import "aos/dist/aos.css";

import InputEntry from "./components/main/InputEntry/page";
import Toast from "@/app/components/main/toast";
import OutputPage from "./components/main/Output/page";
import { OutputContextProvider } from "@/Context/OutputContext";
export default function Home() {
  useEffect(() => {
    AOS.init({duration:1200})
  })

  return (
    <main className='h-full w-full'>
      <div className='flex flex-col gap-20'>
        <Hero />
        <InputEntry></InputEntry>
        
        {/* Hold the Output Data from backend */}
        <OutputContextProvider>
          <OutputPage></OutputPage>
        </OutputContextProvider>
        <Toast></Toast>
        <About />
      </div>
    </main>

    // {/* <main className="flex min-h-screen flex-col items-center p-24 bg-white text-black">
    // </main> */}
  );
}
