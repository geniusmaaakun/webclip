import { parseJsonText } from "typescript";
import { Markdown} from "../types/api/Markdown";
import { debug } from "console";

describe("型", () => {
    it("should be able to parse markdown", () => {
        const jsonData = `{"id": 1, "title": "React dummy", "content": "## h2","path": "test/README.md", "srcurl": "http://test.com", "createdAt": "2021-01-01"}`;

        console.info(jsonData);

        const obj = JSON.parse(jsonData);

        console.info(obj);

        let md: Markdown = {
            id: obj.id,
            title: obj.title,
            path: obj.path,
            srcurl: obj.srcurl,
            content: obj.content,
            created_at: obj.createdAt
        };

        expect(md.id).toBe(1);
        expect(md.title).toBe("React dummy");
        expect(md.content).toBe("## h2");
        expect(md.path).toBe("test/README.md");
        expect(md.srcurl).toBe("http://test.com");
        expect(md.created_at).toBe("2021-01-01");
    });
});