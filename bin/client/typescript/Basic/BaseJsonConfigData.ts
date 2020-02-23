export class BaseJsonConfigData
{
    public static ToInt(value: any): number
    {
        if(typeof(value) == "number")
        {
            return value;
        }
        if(typeof(value) == "string")
        {
            return parseInt(value);
        }
        return 0;
    }

    private m_json: any = null;
    
    private m_valueMap: {[key: string]: any} = {};

    public parseData(json: any)
    {
        this.m_json = json;
        this.onParseData();
        this.m_json = null;
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

    protected getSrcValue(key: string): any
    {
        if(!this.m_json) return null;
        return this.m_json[key];
    }

    protected readList(value, elementReader, target=this): any
    {
        let list = [];
        let valueList = value.split(",");
        let count = valueList.length;
        for(let i = 0; i < count; i++)
        {
            list.push(elementReader.call(target, valueList[i]));
        }
        return list;
    }

    protected readInt(value: any): number
    {
        return BaseJsonConfigData.ToInt(value);
    }

    protected readInt2(value: string): any
    {
        let list = value.split("|");
        let result = {
            x: list.length > 0 ? BaseJsonConfigData.ToInt(list[0]) : 0,
            y: list.length > 1 ? BaseJsonConfigData.ToInt(list[1]) : 0,
        };
        return result;
    }

    protected readInt3(value: string): any
    {
        let list = value.split("|");
        let result = {
            x: list.length > 0 ? BaseJsonConfigData.ToInt(list[0]) : 0,
            y: list.length > 1 ? BaseJsonConfigData.ToInt(list[1]) : 0,
            z: list.length > 2 ? BaseJsonConfigData.ToInt(list[2]) : 0,
        };
        return result;
    }

    protected readInt4(value: string): any
    {
        let list = value.split("|");
        let result = {
            x: list.length > 0 ? BaseJsonConfigData.ToInt(list[0]) : 0,
            y: list.length > 1 ? BaseJsonConfigData.ToInt(list[1]) : 0,
            z: list.length > 2 ? BaseJsonConfigData.ToInt(list[2]) : 0,
            w: list.length > 3 ? BaseJsonConfigData.ToInt(list[3]) : 0,
        };
        return result;
    }

    protected readInt5(value: string): any
    {
        let list = value.split("|");
        let result = {
            x: list.length > 0 ? BaseJsonConfigData.ToInt(list[0]) : 0,
            y: list.length > 1 ? BaseJsonConfigData.ToInt(list[1]) : 0,
            z: list.length > 2 ? BaseJsonConfigData.ToInt(list[2]) : 0,
            w: list.length > 3 ? BaseJsonConfigData.ToInt(list[3]) : 0,
            v: list.length > 4 ? BaseJsonConfigData.ToInt(list[4]) : 0,
        };
        return result;
    }

    protected readString(value: string): string
    {
        return value;
    }
}