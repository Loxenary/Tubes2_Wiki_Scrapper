"use client";
import { showToast } from "@/components/toast";
import {
  SearchWikiInterface,
  useWikiSearchContext,
} from "@/Context/SearchContext";
import { useContext, useState, ChangeEvent, FormEvent } from "react";
import { OutputContext } from "./page";
const EntryWiki = () => {
  const [formValue, setFormValue] = useState({
    FROM: "",
    TO: "",
  });
  const { setData }: SearchWikiInterface = useWikiSearchContext();
  const SearchContext = useContext(OutputContext);

  if (!SearchContext) {
    showToast("Context not found", "error");
    return null;
  }
  const { setOutputState } = SearchContext;

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { id, value } = event.target;
    //TODO: handle autocomplete from wikipedia api
    setFormValue((prevState) => ({
      ...prevState,
      [id]: value,
    }));
  };

  const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    //TODO: implement api request to handle backend
    if (formValue.FROM === formValue.TO) {
      showToast("The Data of From and To are the Same", "warning");
      return;
    } else if (formValue.FROM.length === 0 || formValue.TO.length === 0) {
      showToast("Please fill all the input fields", "error");
      return;
    }
    setData(formValue.FROM, formValue.TO);
    setOutputState(true);
  };
  return (
    <form action="submit" onSubmit={handleSubmit}>
      <div className="w-full justify-center flex gap-x-20">
        <div className="flex flex-col justify-center">
          <label htmlFor="FROM">FROM</label>
          <input
            type="text"
            id="FROM"
            className="bg-gray-100"
            placeholder="wikipedia title"
            value={formValue.FROM}
            onChange={handleChange}
          />
        </div>
        <div className="flex flex-col">
          <label htmlFor="TO">TO</label>
          <input
            type="text"
            id="TO"
            className="bg-gray-100"
            placeholder="wikipedia title"
            value={formValue.TO}
            onChange={handleChange}
          />
        </div>
      </div>
      <div className="flex justify-center items-center w-full my-12">
        <button className="text-white bg-blue-400 w-20 h-10 text-xl">
          FIND
        </button>
      </div>
    </form>
  );
};

export default EntryWiki;
