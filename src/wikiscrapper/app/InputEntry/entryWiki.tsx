import { useState, createContext, useContext } from "react";

const EntryWiki = () => {
  // const [fromValue, setFromValue] = useState("");
  // const [toValue, setToValue] = useState("");

  // const handleFromSubmit = (event) => {
  //   event.preventDefault();
  //   // Handle form submission for FROM here
  //   console.log("Submitted FROM:", fromValue);
  // };

  // const handleToSubmit = (event) => {
  //   event.preventDefault();
  //   // Handle form submission for TO here
  //   console.log("Submitted TO:", toValue);
  // };

  return (
    <form action="submit">
      <div className="w-full justify-center flex gap-x-20">
        <div className="flex flex-col justify-center">
          <label htmlFor="FROM">FROM</label>
          <input
            type="text"
            id="FROM"
            className="bg-gray-100"
            placeholder="wikipedia title"
          />
        </div>
        <div className="flex flex-col">
          <label htmlFor="TO">TO</label>
          <input
            type="text"
            id="TO"
            className="bg-gray-100"
            placeholder="wikipedia title"
          />
        </div>
      </div>
    </form>
  );
};

export default EntryWiki;
