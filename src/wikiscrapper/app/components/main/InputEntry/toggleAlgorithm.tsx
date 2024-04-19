"use client";
import ToggleSlider from "@/app/components/main/toggleSlider";
import DefaultImage from "@/public/switch-body-light.png";
import ToggledImage from "@/public/switch-body-dark.png";
import {SearchWikiInterface,useWikiSearchContext } from "@/Context/SearchContext";

const ToggleAlgorithm = () => {
  const { setAlgorithm }: SearchWikiInterface = useWikiSearchContext();
  const OnSliderChange = (state: boolean) => {
    if (state) {
      setAlgorithm("IDS");
    } else {
      setAlgorithm("BFS");
    }
  };
  return (
    <div className="w-full flex justify-center items-center my-10 gap-x-10" data-aos="fade-left">
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
