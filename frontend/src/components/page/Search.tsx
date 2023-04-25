import React, {useState, useEffect} from "react"
import axios from 'axios';

interface Markdown {
    title: string
    content: string
    path: string
    srcUrl: string
}

export const Search = () => {
    //文字列を受け取るためのstate
    const [input, setInput] = useState('');
    //マスターデータ
    const [data, setData] = useState<Markdown[]>([]);
    //表示するデータ
    const [resultData, setResultData] = useState<Markdown[]>([]);

  
    //入力されたら、inputに格納
    const handleChange = (e: React.ChangeEvent<{ value: string }>) => {
      setInput(e.target.value);
      search(e.target.value);
    };

    useEffect(() => {
        const fetchData = async () => {
            const result = await axios(
                'http://localhost:8080/api/markdowns',
            );
            if (result.data) {
                console.log(result.data);
                setData(result.data);
                setResultData(result.data);
            }
        };
        console.log("fetchData");
        fetchData();
    }, []);

    // 検索欄への入力値での絞り込み
    const search = (value :string) => {
    if (value === '') {
        setResultData(data);
        return;
    }
    const searchKeywords  = value.trim()
    .toLowerCase()
    .match(/[^\s]+/g);
    // 検索欄への入力が空の場合は早期return
    if (searchKeywords === null) {
        setResultData(data);
      return;
    }

    const serchedPosts = data.filter(
      (data) => {
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
            return searchKeywords.every((kw) => {
                if (data.title !== undefined && data.title !== null && data.title.toString().toUpperCase().indexOf(kw.toString().toUpperCase()) !== -1) {
                    return true;
                }
                return false;
            });
        }
    );

    console.log(serchedPosts);

    setResultData(serchedPosts);
  }
    
  
    return (
      <div>
        <h1>WebClip</h1>
        <input
          type="text"
          value={input}
          onChange={handleChange}
          //onKeyPress={handleKeyPress}
          placeholder="検索キーワードを入力"
        />
        <ul>
          {resultData.map((item, index) => (
            <li key={index}>{item.title}</li> // APIレスポンスに応じて、適切なデータを表示してください
          ))}
        </ul>
      </div>
    );
}
