import Title from "./title";
import ToggleAlgorithm from "./toggleAlgorithm";
import EntryWiki from "./entryWiki";
import FindButton from "./findButton";
const InputEntry = () => {
  return (
    <div className="w-full justify-center items-center flex flex-col gap-y-5">
      <Title></Title>
      <ToggleAlgorithm></ToggleAlgorithm>
      <EntryWiki></EntryWiki>
      <FindButton></FindButton>
    </div>
  );
};

export default InputEntry;
