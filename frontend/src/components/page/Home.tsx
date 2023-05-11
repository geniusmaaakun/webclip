import { MarkdownEditor } from "./Markdown";
import { Search } from "./Search";

export const Home = () => {
    //これは不要
    // const { markdowns } = useMarkdowns();
    
    return (
        <>
            <div className="left">
                <Search />
            </div>
            <div className="right">
                <MarkdownEditor  />
            </div>
        </>
    )
}
export default Home;