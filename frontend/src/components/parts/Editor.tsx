import React, { useEffect, useState } from "react";
import { Markdown } from "../../types/api/Markdown";
import SimpleMde from "react-simplemde-editor";
import DOMPurify from "dompurify";
import { marked } from "marked";
import "easymde/dist/easymde.min.css";
//npm install --save @types/dompurifyをしないといけない
import highlightjs from "highlight.js";
import "highlight.js/styles/github.css";
import { useLoadMarkdown } from "../../hooks/markdowns/useMarkdowns";
import { useMarkdowns } from "../../hooks/providers/useMarkdownsProvider";
import { AxiosRequestConfig } from "axios";

interface Props {
  id: string;
}

export const Editor = (props: Props) => {
  const { id } = props;
  // ハイライトの設定
  marked.setOptions({
    highlight: (code, lang) => {
      return highlightjs.highlightAuto(code, [lang]).value;
    },
  });

  const [markdownValue, setMarkdownValue] = useState<string>("");

  const { loadMarkdown } = useLoadMarkdown();

  //?
  const { getMarkdownById } = useMarkdowns();
  console.log(id);

  //使ってない
  const markdown = getMarkdownById(id!);
  console.log("mds", markdown);

  // if (markdown) {
  //   setMarkdownValue(markdown!.content || "");
  // }

  //idからAPIを叩いて、データを取得する
  //取得したデータをvalueに入れる
  useEffect(() => {
    //API通信を中断するための処理
    const controller = new AbortController();

    //APIを叩く処理
    //取得したデータをvalueに入れる

    async function fetchData() {
      const options: AxiosRequestConfig = {
        signal: controller.signal, //AbortControllerとAxiosの紐付け
      };
      const md = await loadMarkdown(id!, options);

      //console.log(markdown!.content);
      setMarkdownValue(md ? md.content : "");
    }

    if (id && markdown?.content === undefined) {
      console.log("fetch!!");
      fetchData();
    }

    return () => {
      //画面遷移する時(アンマウントする時)に通信を中止する
      controller.abort();
    };
  }, [id]);

  const onChange = (value: string) => {
    setMarkdownValue(value);
    //ファイルを保存する処理
  };
  return (
    <>
      <SimpleMde value={markdownValue} onChange={onChange} />
      <div
        dangerouslySetInnerHTML={{
          __html: DOMPurify.sanitize(marked(markdownValue)),
        }}
      ></div>
    </>
  );
};
