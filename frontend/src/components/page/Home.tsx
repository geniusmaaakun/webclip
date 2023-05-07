import MarkdownEditor from "./Markdown";
import { Search } from "./Search";
import { useMarkdowns } from "../../hooks/providers/useMarkdownsProvider";
import { useLoadMarkdowns } from "../../hooks/markdowns/useMarkdowns";


export const Home = () => {
    const { markdowns } = useMarkdowns();
    
    return (
        <>
            <div className="left">
                <Search markdowns={markdowns} />
            </div>
            <div className="right">
                <MarkdownEditor markdowns={markdowns} />
            </div>
        </>
    )
}
export default Home;