import OutputTitle from "./title";
import ArticlesOut from "./articlesOut";
import RouteOutput from "./Route";
const OutputPage = () => {
  return (
    <div className="flex flex-col mx-20 my-10 gap-10" data-aos="fade-down">
      <OutputTitle />
      <ArticlesOut></ArticlesOut>
      <RouteOutput></RouteOutput>
    </div>
  );
};

export default OutputPage;
