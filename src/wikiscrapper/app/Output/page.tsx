import OutputTitle from "./title";
import ArticlesOut from "./articlesOut";
import RouteOutput from "./Route";
import { IOutputContext } from "../InputEntry/page";
const OutputPage = () => {
  return (
    <div className="flex flex-col my-10 gap-10">
      <OutputTitle />
      <ArticlesOut></ArticlesOut>
      <RouteOutput></RouteOutput>
    </div>
  );
};

export default OutputPage;
