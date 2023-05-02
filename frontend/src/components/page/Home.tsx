import MarkdownEditor from "./Markdown";
import { Search } from "./Search";

export const Home = () => {
    return (
        <>
            <div className="left">
                <Search />
            </div>
            <div className="right">
                <MarkdownEditor />
            </div>
        </>
    )
}
export default Home;