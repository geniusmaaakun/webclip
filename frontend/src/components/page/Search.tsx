import React, { useState, useEffect } from "react";
import { useMarkdowns } from "../../hooks/providers/useMarkdownsProvider";
import axios from "axios";
import { Markdown } from "../../types/api/Markdown";
import { useLoadMarkdowns } from "../../hooks/markdowns/useMarkdowns";

import { useNavigate } from 'react-router-dom'; //v6
/*
interface Markdown {
  title: string;
  content: string;
  path: string;
  srcUrl: string;
}
*/

type Props = {
  markdowns: Markdown[];
};

export const Search = React.memo(() => {
  //文字列を受け取るためのstate
  const [input, setInput] = useState("");
  //マスターデータ
  const [data, setData] = useState<Markdown[]>([]);
  const { markdowns, setMarkdowns } = useMarkdowns();
  const { loadMarkdowns } = useLoadMarkdowns();

  //表示するデータ
  const [resultData, setResultData] = useState<Markdown[]>([]);

  const navigate = useNavigate();

  //入力されたら、inputに格納
  const handleChange = (e: React.ChangeEvent<{ value: string }>) => {
    setInput(e.target.value);
    // let searchedmd =  SearchMarkdown(e.target.value);
    // console.log(searchedmd);
    // setResultData(searchedmd);

    search(e.target.value);
  };

  useEffect(() => {   
    async function fetchData() {
      const mds = await loadMarkdowns();
      console.log(mds);
      setResultData(mds);
    }

    //setResultData(markdowns);

    fetchData();
  }, [loadMarkdowns]);

    // 検索欄への入力値での絞り込み
    //React.memo
  const search = (value: string) => {
    if (value === "") {
      setResultData(markdowns);
      return;
    }
    //　match  正規表現にマッチするか？gオプションで配列を返すない場合は、一つだけ
    const searchKeywords = value
      .trim()
      .toLowerCase()
      .match(/[^\s]+/g);
    // 検索欄への入力が空の場合は早期return
    if (searchKeywords === null) {
      setResultData(markdowns);
      return;
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

    console.log(serchedPosts);

    setResultData(serchedPosts);
  };

  const onClickMarkdown = (id: string) => {
    navigate(`/markdowns/${id}`);
  }

  const onClickHome = () => {
    navigate(`/`);
  }

  return (
    <div>
      <h1 onClick={onClickHome}>WebClip</h1>
      <input
        type="text"
        value={input}
        onChange={handleChange}
        //onKeyPress={handleKeyPress}
        placeholder="検索キーワードを入力"
      />
      <ul>
        {resultData && resultData.map((item, index) => (
          //詳細ページを作成
          <li key={index} onClick={() => onClickMarkdown(item.id)}>
            {item.title}
          </li> // APIレスポンスに応じて、適切なデータを表示してください
        ))}
      </ul>
    </div>
  );
});
