"use client";
import ToggleSlider from "@/app/components/main/toggleSlider";
import ToggleImage from "@/public/image/grad1.png";
import DefaultImage from "@/public/image/grad2.png";
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
          toggled: ToggleImage,
        }}
        onChange={OnSliderChange}
      ></ToggleSlider>
      <h1>IDS</h1>
    </div>
  );
};

export default ToggleAlgorithm;
