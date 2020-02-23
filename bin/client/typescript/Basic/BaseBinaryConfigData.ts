export class BaseBinaryConfigData
{
    private m_binary: Uint8Array = null;
    private m_index: number = 0;

    private m_valueMap: {[key: string]: any} = {};

    public parseData(binary: Uint8Array, fromIndex: number = 0): number
    {
        this.m_binary = binary;
        this.m_index = fromIndex;
        this.onParseData();
        this.m_binary = null;
        return this.m_index;
    }

    protected onParseData()
    {

    }

    protected setConfigValue(key: string, value: any)
    {
        this.m_valueMap[key] = value;
    }

    public getConfigValue(key: string): any
    {
        return this.m_valueMap[key];
    }

    public hasConfigValue(key: string): boolean
    {
        if(typeof(this.m_valueMap[key]) == "undefined")
        {
            return false;
        }
        return true;
    }

    protected readList(elementReader, target=this): any
    {
        let list = [];
        let count = this.readInt();
        for(let i = 0; i < count; i++)
        {
            list.push(elementReader.call(target));
        }
        return list;
    }

    protected readInt(): number
    {
        let b1: number = this.m_binary[this.m_index + 3];
        let b2: number = this.m_binary[this.m_index + 2];
        let b3: number = this.m_binary[this.m_index + 1];
        let b4: number = this.m_binary[this.m_index + 0];
        this.m_index += 4;
        let result = (b1 << (3 * 8)) + (b2 << (2 * 8)) + (b3 << (1 * 8)) + b4;
        return result;
    }

    protected readInt2(): any
    {
        let result = {
            x: this.readInt(),
            y: this.readInt(),
        };
        return result;
    }

    protected readInt3(): any
    {
        let result = {
            x: this.readInt(),
            y: this.readInt(),
            z: this.readInt(),
        };
        return result;
    }

    protected readInt4(): any
    {
        let result = {
            x: this.readInt(),
            y: this.readInt(),
            z: this.readInt(),
            w: this.readInt(),
        };
        return result;
    }

    protected readInt5(): any
    {
        let result = {
            x: this.readInt(),
            y: this.readInt(),
            z: this.readInt(),
            w: this.readInt(),
            v: this.readInt(),
        };
        return result;
    }

    protected readString(): string
    {
        let len = this.readInt();
        let bytes = this.readBytes(len);
        let str = "";
        let count = bytes.length;
        for(let i = 0; i < count; i++)
        {
            str += String.fromCharCode(bytes[i]);
        }
        let index = 0;
        let head = null;
        let body = null;
        let tail = null;
        count = str.length;
        for(let i = 0; i < count; i++)
        {
            index = str.indexOf("\\u");
            if(index == -1) break;
            head = index ? str.substring(0, index) : null;
            body = str.substring(index + 2, index + 2 + 4);
            tail = str.substring(index + 2 + 4);
            if(head)
            {
                str = head + String.fromCharCode(parseInt(body, 16)) + tail;
            }
            else
            {
                str = String.fromCharCode(parseInt(body, 16)) + tail;
            }
        }
        return str;
    }

    protected readBytes(len: number): number[]
    {
        let bytes: number[] = [];
        for(let i = 0; i < len; i++)
        {
            if(this.m_index >= this.m_binary.length)
            {
                bytes.push(0);
            }
            else
            {
                bytes.push(this.m_binary[this.m_index]);
                this.m_index++;
            }
        }
        return bytes;
    }
}