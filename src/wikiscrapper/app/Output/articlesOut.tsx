const ArticlesOut = () => {
  return (
    <div className="justify-between items-center text-2xl flex gap-x-20">
      <div className="flex gap-5 justify-center items-center">
        <h1>Articles Checked</h1>
        <div className="bg-gray-500 w-10 h-10"></div>
      </div>
      <div className="flex gap-5 justify-center items-center">
        <h1>Articles Passed</h1>
        <div className="bg-gray-500 w-10 h-10"></div>
      </div>
    </div>
  );
};

export default ArticlesOut;
