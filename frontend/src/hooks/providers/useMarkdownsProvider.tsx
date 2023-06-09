import React, {
    createContext,
    Dispatch,
    ReactNode,
    SetStateAction,
    useContext,
    useState,
    useEffect
} from "react";

import { Markdown } from "../../types/api/Markdown";

type MarkdownsContextType = {
    markdowns: Markdown[];
    setMarkdowns: Dispatch<SetStateAction<Markdown[]>>;
    getMarkdownById: (id: string) => Markdown | undefined;
};

const MarkdownsContext = createContext<MarkdownsContextType>({} as MarkdownsContextType);

//routerでこのコンポーネントをラップすることで、stateを共有できる
export const MarkdownProvider = (props: { children: ReactNode }) => {
    const { children } = props;
    const [markdowns, setMarkdowns] = useState<Markdown[]>([]);

    //ここでAPIでアクセスしておく？
    useEffect(() => {

    }, [])

    const getMarkdownById = (id: string): Markdown | undefined => {
        //console.log(markdowns[0].id)
        //console.log(id)
        // console.log(typeof id)
        //console.log(typeof markdowns[0].id ? markdowns[0].id : "number")

        //console.log("1" === id)
        //console.log(id, "1" ===  id)

        //console.log(markdowns[0].id === id)
        let finedMarkdown = markdowns.find(markdown => {
            //何故か型が違うので、Stringでキャストする
            //console.log(typeof id)
            //console.log(typeof markdown.id)
            //console.log(String(markdown.id) === id)
            return String(markdown.id) === id
        })
        //console.log(finedMarkdown)
        return finedMarkdown
    }

    //実際に提供されるもの
    return (
        <MarkdownsContext.Provider value={{ markdowns, setMarkdowns, getMarkdownById }}>
            {children}
        </MarkdownsContext.Provider>
    );
};

export const useMarkdowns = (): MarkdownsContextType => useContext(MarkdownsContext);