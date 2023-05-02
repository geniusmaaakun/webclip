import React, {
  useEffect,
  useState,
  useCallback,
  useMemo,
  ChangeEvent,
} from "react";
import { useMarkdowns } from "../providers/useMarkdownsProvider";

import api from "../../services/api";
import { Markdown } from "../../types/api/Markdown";

export const useLoadMarkdown = () => {
  const { markdowns, setMarkdowns } = useMarkdowns();

  const loadMarkdown = useCallback(
    (id: string) => {
      const fetchMarkdown = async () => {
        const res = await fetch(`http://localhost:8080/api/markdowns/${id}`);
        const json = await res.json();
        setMarkdowns(
          markdowns.map((markdown) => {
            //idが一致したら、jsonに置き換える
            if (markdown.id === json.id) {
              return json;
            }
            //idが一致しなかったら、そのまま返す
            return markdown;
          })
        );
      };
      fetchMarkdown();
      console.log(markdowns)
    },
    [setMarkdowns]
  );

  return { loadMarkdown };
};

export const useLoadMarkdowns = () => {
  const { markdowns, setMarkdowns } = useMarkdowns();
  let fetchMarkdowns :Markdown[] = []
  //useCallbackで関数をメモ化する
  //useCallbackはReact.memoと併用する。
  //React.memoの使用例 https://qiita.com/seira/items/8a170cc950241a8fdb23
  const loadMarkdowns = useCallback(async() => {
    /*
    //非同期関数 awaitで待機する
    const fetchMarkdowns = async () => {
      //  test用
      const res = await fetch("http://localhost:8080/api/markdowns");
      // build用
      //const res = await fetch("http://localhost:8080/api/markdowns/");
      const json = await res.json();

      console.log(res)
     
      console.log(json)

    //どちらでもいい
      //setMarkdowns(json);
      setMarkdowns(
            markdowns.map((markdown) => {
            //idが一致したら、jsonに置き換える
            if (markdown.id === json.id) {
                return json;
            }
            //idが一致しなかったら、そのまま返す
            return markdown;
            })
      );
    };
    */
    //fetchMarkdowns();

    await api.get("/api/markdowns").then((res) => {
        console.log(res.data)
        let resMarkdowns: Markdown[] = res.data;

        /*
        console.log(resMarkdowns)
        resMarkdowns.forEach((resMarkdown) => {
            console.log(resMarkdown)
        });
        */
        fetchMarkdowns = resMarkdowns.map((markdown) => {
            let md : Markdown = {
                id: markdown.id,
                title: markdown.title,
                content: markdown.content,
                path: markdown.path,
                srcurl: markdown.srcurl,
                created_at: markdown.created_at,    
            }
            return md ;
        });
        setMarkdowns(fetchMarkdowns);
    })
    // console.log("test")
    // console.log(markdowns)
    return fetchMarkdowns;
  }, [setMarkdowns]); //この関数を使ったコンポーネントが再レンダリングされるのを防ぐ

  return { loadMarkdowns };
};


export const useSearchMarkdowns = () => {
    // 検索欄への入力値での絞り込み
    const SearchMarkdown = (value: string) => {
        const { markdowns, setMarkdowns } = useMarkdowns();

        if (value === "") {
            return markdowns;
        }
        //　match  正規表現にマッチするか？gオプションで配列を返すない場合は、一つだけ
        const searchKeywords = value
            .trim()
            .toLowerCase()
            .match(/[^\s]+/g);
        // 検索欄への入力が空の場合は早期return
        if (searchKeywords === null) {
            return markdowns; 
        }

        const serchedPosts = markdowns.filter((data) => {
            //各valueに対して、入力値と一致するかを判定
            /*
            Object.values(post).filter(
                (item: string) =>
                item !== undefined &&
                item !== null &&
                item.toString().toUpperCase().indexOf(value.toString().toUpperCase()) !== -1
            ).length > 0
            */

            //タイトルのみ検索
            // if (post.title !== undefined && post.title !== null && post.title.toString().toUpperCase().indexOf(value.toString().toUpperCase()) !== -1) {
            //         return true;
            // }

            //キーワードを全て含むか
            return searchKeywords.every((kw) => {
            if (
                data.title !== undefined &&
                data.title !== null &&
                data.title
                .toString()
                .toUpperCase()
                .indexOf(kw.toString().toUpperCase()) !== -1
            ) {
                return true;
            }
            return false;
            });
        });
        return serchedPosts;
    }

    return { SearchMarkdown };
}