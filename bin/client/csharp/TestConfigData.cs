////////////////////////////////////////////////////////////////
//
// 此文件由Configo自动生成
// 请勿随意修改，以免造成不必要的损失
//
////////////////////////////////////////////////////////////////


public class TestConfigData : BaseBinaryConfigData 
{ 
    protected void onParseData()
    {
        setConfigValue("id", readInt());// ColumnIndex: 0
        setConfigValue("val_str", readString());// ColumnIndex: 1
        setConfigValue("val_int", readInt());// ColumnIndex: 2
        setConfigValue("val_int_list", readList(this.readInt)());// ColumnIndex: 3
        setConfigValue("int_two", readInt2());// ColumnIndex: 4
        setConfigValue("int_tri", readInt3());// ColumnIndex: 5
        setConfigValue("int_two_list", readList(this.readInt2)());// ColumnIndex: 6
    }

    public int getId()
    {
        return getConfigValue<int>("id");
    }

    public string getValStr()
    {
        return getConfigValue<string>("val_str");
    }

    public int getValInt()
    {
        return getConfigValue<int>("val_int");
    }

    public List<int> getValIntList()
    {
        return getConfigValue<List<int>>("val_int_list");
    }

    public Int2 getIntTwo()
    {
        return getConfigValue<Int2>("int_two");
    }

    public Int3 getIntTri()
    {
        return getConfigValue<Int3>("int_tri");
    }

    public List<Int2> getIntTwoList()
    {
        return getConfigValue<List<Int2>>("int_two_list");
    }
} 
