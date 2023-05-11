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
import { AxiosRequestConfig } from "axios";

//プロバイダにまとめる？

export const useLoadMarkdown = () => {
  const { markdowns, setMarkdowns } = useMarkdowns();
  let md: Markdown;

  const loadMarkdown = useCallback(
    async (id: string, config?: AxiosRequestConfig) => {
      /*  
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
        return json;
      };
      const md = await fetchMarkdown();
      //console.log(markdowns)
      return md;
      */
        await api.get(`/api/markdowns/${id}`, config).then((res) => {
            console.log(res.data)

            md = {
                id: res.data.id,
                title: res.data.title,
                content: res.data.content,
                path: res.data.path,
                srcurl: res.data.srcurl,
                created_at: res.data.created_at,
            };
            //アロー関数で、前回のデータを元に新しいデータを更新することで、更新が保証される。
            setMarkdowns((prevMarkdowns) =>
                prevMarkdowns.map((markdown) => {
                //idが一致したら、jsonに置き換える
                if (markdown.id === String(md.id)) {
                    return md;
                }
                //idが一致しなかったら、そのまま返す
                return markdown;
                })
            );
        }).catch((err) => {
            console.log(err);
        });

        return md;
    },
    [setMarkdowns]
  );

  return { loadMarkdown };
};

export const useLoadMarkdowns = () => {
  const { markdowns, setMarkdowns } = useMarkdowns();
  //setMarkdowns(markdowns); が非同期な為、setMarkdowns が完了していない可能性がある。
  let fetchMarkdowns :Markdown[] = []
  //useCallbackで関数をメモ化する
  //useCallbackはReact.memoと併用する。
  //React.memoの使用例 https://qiita.com/seira/items/8a170cc950241a8fdb23
  const loadMarkdowns = useCallback(async(config?: AxiosRequestConfig) => {
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

    await api.get("/api/markdowns", config).then((res) => {
        //console.log(res.data)
        let resMarkdowns: Markdown[] = res.data;

        /*
        console.log(resMarkdowns)
        resMarkdowns.forEach((resMarkdown) => {
            console.log(resMarkdown)
        });
        */
       console.log(res.data)

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
        //console.log(markdowns)
    }).catch((err) => {
        console.log(err);
    })

     //console.log("test")
    // console.log(markdowns)
    return fetchMarkdowns;
  }, [setMarkdowns]); //この関数を使ったコンポーネントが再レンダリングされるのを防ぐ

  return { loadMarkdowns };
};

