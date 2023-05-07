// import React, { useState } from "react";
// import SimpleMde from "react-simplemde-editor";
// import "easymde/dist/easymde.min.css";

//基本的なエディタの実装
// export const MarkdownEditor = () => {
//   const [markdownValue, setMarkdownValue] = useState("Initial value");

//   const onChange = (value) => {
//     setMarkdownValue(value);
//   };

//   return <SimpleMde value={markdownValue} onChange={onChange} />;
// };

// export default MarkdownEditor;


// import React, { useState } from "react";
// import SimpleMde from "react-simplemde-editor";
// import "easymde/dist/easymde.min.css";
// import {marked} from "marked";
// import DOMPurify from "dompurify";
// import highlightjs from "highlight.js";
// import "highlight.js/styles/github.css";

// //プレビュー機能の追加
// const MarkdownEditor = () => {
//     const [markdownValue, setMarkdownValue] = useState("Initial value");
   
//     const onChange = (value) => {
//       setMarkdownValue(value);
//     };
   
//     return (
//     <>
//        <SimpleMde value={markdownValue} onChange={onChange} />
//        <div>
//           <div dangerouslySetInnerHTML={{__html: DOMPurify.sanitize(marked(markdownValue))}}></div>
//        </div>
//     </>
//     );
//    };
   
// export default MarkdownEditor;


import React, { useEffect, useState } from "react";
import SimpleMde from "react-simplemde-editor";
import "easymde/dist/easymde.min.css";
import {marked} from "marked";
//npm install --save @types/dompurifyをしないといけない
import DOMPurify from "dompurify";
import highlightjs from "highlight.js";
import "highlight.js/styles/github.css";
//ハイライトをつけよう

import {useParams} from "react-router-dom"
import { useLoadMarkdown } from "../../hooks/markdowns/useMarkdowns";
import { useMarkdowns} from "../../hooks/providers/useMarkdownsProvider";
import { Markdown } from "../../../src/types/api/Markdown"

/**
 ```js
const [test, setTest] = useState();
```
```go
sl := 1
```
 */

type Props = {
  markdowns: Markdown[];
};

const MarkdownEditor = React.memo(() => {
  //const { markdowns } = props;

    // ハイライトの設定
  marked.setOptions({
    highlight: (code, lang) => {
      return highlightjs.highlightAuto(code, [lang]).value;
    },
  });

  //urlからidを取得する
  const [markdownValue, setMarkdownValue] = useState("");
  const {id} = useParams();
  const { loadMarkdown } = useLoadMarkdown();

  //?
  const { getMarkdownById } = useMarkdowns();
  // console.log(id);

  //使ってない
  const markdown = getMarkdownById(id!);
  console.log("mds", markdown)

  // if (markdown) {
  //   setMarkdownValue(markdown!.content || "");
  // }
  
  //idからAPIを叩いて、データを取得する
  //取得したデータをvalueに入れる
  useEffect(() => {
    //APIを叩く処理
    //取得したデータをvalueに入れる
    
    async function fetchData() {
      const md = await loadMarkdown(id!);
      
      //console.log(markdown!.content);
      setMarkdownValue(md.content || "");
    }

    if (id && markdown?.content === undefined) {
      console.log("fetch!!");
      fetchData();
    }
  }, [id]);
  
 
  const onChange = (value: string) => {
    setMarkdownValue(value);
    //ファイルを保存する処理
  };
 
  if (!id) {
    return null;
  }
  return (
    <div>
      <SimpleMde value={markdownValue} onChange={onChange} />
      <div
        dangerouslySetInnerHTML={{
          __html: DOMPurify.sanitize(marked(markdownValue)),
        }}
      ></div>
    </div>
  );
 });
 
 export default MarkdownEditor;
 
