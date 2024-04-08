"use client";
import ToggleSlider from "@/components/toggleSlider";
import DefaultImage from "@/public/switch-body-light.png";
import ToggledImage from "@/public/switch-body-dark.png";
import { useState } from "react";
const ToggleAlgorithm = () => {
  const [CurrentAlgorithm, SetCurrentAlgorithm] = useState("IDS");

  const OnSliderChange = (state: boolean) => {
    if (state) {
      SetCurrentAlgorithm("BFS");
    } else {
      SetCurrentAlgorithm("IDS");
    }
  };
  return (
    <div className="w-full flex justify-center items-center gap-x-10 ">
      {/* Width and Height is on pixel */}
      <h1>BFS</h1>
      <ToggleSlider
        height={40}
        width={120}
        background={{
          default: DefaultImage,
          toggled: ToggledImage,
        }}
        onChange={OnSliderChange}
      ></ToggleSlider>
      <h1>IDS</h1>
    </div>
  );
};

export default ToggleAlgorithm;
