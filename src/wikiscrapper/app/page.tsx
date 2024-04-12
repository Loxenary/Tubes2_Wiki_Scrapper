"use client";

import Image from "next/image";
import Head from 'next/head';
import Hero from "./components/main/Hero";
import About from "./components/main/About";
import { useEffect } from "react";

import AOS from "aos";
import "aos/dist/aos.css";

export default function Home() {
  useEffect(() => {
    AOS.init({duration:1200})
  })

  return (
    <main className='h-full w-full'>
      <div className='flex flex-col gap-20'>
        <Hero />
        <About />
      </div>
    </main>
  );
}
