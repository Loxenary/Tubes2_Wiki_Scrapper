"use client";
import { showToast } from "@/components/toast";
import {
  SearchWikiInterface,
  useWikiSearchContext,
} from "@/Context/SearchContext";
import { useContext, useState, ChangeEvent, FormEvent } from "react";
import { BoolOutputSetup } from "./page";
import { IOutputContext } from "../Output/outputData";
import { useOutputContext } from "@/Context/OutputContext";
import { LoadingBar } from "@/components/loading";
const EntryWiki = () => {
  const [formValue, setFormValue] = useState({
    FROM: "",
    TO: "",
  });

  const [isLoading, setisLoading] = useState(false);

  const { setOutputData }: IOutputContext = useOutputContext();

  const { setDataInput, Algorithm }: SearchWikiInterface =
    useWikiSearchContext();
  const SearchContext = useContext(BoolOutputSetup);

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

  function handleOutputData(data: any) {
    const { checkcount, listPath, numpassed, time } = data;
    setOutputData(checkcount, numpassed, time, listPath);
  }

  function handleInputData() {
    setDataInput(formValue.FROM, formValue.TO);
    const dataUsed = {
      ...formValue,
      algorithm: Algorithm,
    };
    setFormValue(dataUsed);
  }

  function handleInputEdgeCases() {
    if (formValue.FROM === formValue.TO) {
      throw new Error("The Data of From and To are the Same");
    } else if (formValue.FROM.length === 0 || formValue.TO.length === 0) {
      throw new Error("Please fill all the input fields");
    }
  }

  const HandleAPI = async () => {
    const res = await fetch("/api/postData", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(formValue),
    });

    if (!res.ok) {
      throw new Error("failed to fetch");
    }
    const data = await res.json();
    return data;
  };

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    //TODO: implement api request to handle backend
    handleInputData();
    try {
      handleInputEdgeCases();
    } catch (error) {
      showToast(error + "", "error");
      return;
    }
    try {
      setisLoading(true);
      setOutputState(false);
      const data = await HandleAPI();
      setisLoading(false);
      handleOutputData(data);
    } catch (error) {
      showToast("Error :" + error, "error");
      return;
    }
    setOutputState(true);
  };
  return (
    <>
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
      <div>
        <LoadingBar isLoading={isLoading} />
      </div>
    </>
  );
};

export default EntryWiki;
