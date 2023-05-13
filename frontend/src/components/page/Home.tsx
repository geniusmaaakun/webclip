import { MarkdownEditor } from "./Markdown";
import { Search } from "./Search";
import "./home.css"

export const Home = () => {
    //これは不要
    // const { markdowns } = useMarkdowns();
    
    //flexboxで左右に分ける
    return (
        <div className="container">
            <div className="left">
                    <Search />
            </div>
            <div className="right">
                    <MarkdownEditor />
            </div>
        </div>
    )
}
export default Home;