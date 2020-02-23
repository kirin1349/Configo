////////////////////////////////////////////////////////////////
//
// 此文件由Configo自动生成
// 请勿随意修改，以免造成不必要的损失
//
////////////////////////////////////////////////////////////////

import { BaseJsonConfigData } from "./Basic/BaseJsonConfigData";
import { Num2 } from "./Basic/Num2";
import { Num3 } from "./Basic/Num3";

export class TestConfigData extends BaseJsonConfigData 
{ 
    protected onParseData()
    {
        this.setConfigValue("id", this.readInt(this.getSrcValue("id")));// ColumnIndex: 0
        this.setConfigValue("val_str", this.readString(this.getSrcValue("val_str")));// ColumnIndex: 1
        this.setConfigValue("val_int", this.readInt(this.getSrcValue("val_int")));// ColumnIndex: 2
        this.setConfigValue("val_int_list", this.readList(this.getSrcValue("val_int_list"), this.readInt, this));// ColumnIndex: 3
        this.setConfigValue("int_two", this.readInt2(this.getSrcValue("int_two")));// ColumnIndex: 4
        this.setConfigValue("int_tri", this.readInt3(this.getSrcValue("int_tri")));// ColumnIndex: 5
        this.setConfigValue("int_two_list", this.readList(this.getSrcValue("int_two_list"), this.readInt2, this));// ColumnIndex: 6
    }

    getId(): number
    {
        return this.getConfigValue("id");
    }

    getValStr(): string
    {
        return this.getConfigValue("val_str");
    }

    getValInt(): number
    {
        return this.getConfigValue("val_int");
    }

    getValIntListList(): number[]
    {
        return this.getConfigValue("val_int_list");
    }

    getIntTwo(): Num2
    {
        return this.getConfigValue("int_two");
    }

    getIntTri(): Num3
    {
        return this.getConfigValue("int_tri");
    }

    getIntTwoListList(): Num2[]
    {
        return this.getConfigValue("int_two_list");
    }
} 
