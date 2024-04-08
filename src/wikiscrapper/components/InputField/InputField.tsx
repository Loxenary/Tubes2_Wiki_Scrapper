import { useState } from "react";

interface InputProperty {
  title: string;
  textareaproperties?: TextAreaProperty;
  onChange: (value: string) => void;
}

interface TextAreaProperty {
  width: number;
  height: number;
  placeholder?: string;
  defaultValue?: string;
}

const InputField: React.FC<InputProperty> = ({
  title,
  textareaproperties,
  onChange,
}) => {
  const [InputValue, SetInputValue] = useState("");

  const HandleInputValue = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = event.target.value;
    SetInputValue(value);
    onChange(value);
  };

  return (
    <form>
      <label
        htmlFor="search"
        className="mb-2 text-sm font-medium text-gray-900 dark:text-white"
      >
        {title}
      </label>
      <input
        id="search"
        placeholder={textareaproperties?.placeholder}
        className={`w-46 h-12`}
        onChange={HandleInputValue}
        value={InputValue} // Controlled input value
      ></input>
    </form>
  );
};

export default InputField;
