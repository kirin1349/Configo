using System;
using System.Collections.Generic;

public class BaseBinaryConfigData
{
    private byte[] m_binary = null;
    private int m_index = 0;

    private Dictionary<string, object> m_valueMap = new Dictionary<string, object>();

    public void setData(byte[] binary)
    {
        m_binary = binary;
        onParseData();
    }

    protected void onParseData()
    {

    }

    protected void setConfigValue(string key, object value)
    {
        if (m_valueMap.ContainsKey(key))
        {
            m_valueMap[key] = value;
        }
        else
        {
            m_valueMap.Add(key, value);
        }
    }

    protected T getConfigValue<T>(string key)
    {
        if (m_valueMap.ContainsKey(key))
        {
            return (T)m_valueMap[key];
        }
        return default(T);
    }

    protected List<T> readList<T>(Func<T> elementReader)
    {
        List<T> list = new List<T>();
        int count = readInt();
        for (int i = 0; i < count; i++)
        {
            list.Add(elementReader());
        }
        return list;
    }

    protected int readInt()
    {
        int b1 = m_binary[m_index + 3];
        int b2 = m_binary[m_index + 2];
        int b3 = m_binary[m_index + 1];
        int b4 = m_binary[m_index + 0];
        m_index += 4;
        int result = (b1 << (3 * 8)) + (b2 << (2 * 8)) + (b3 << (1 * 8)) + b4;
        return result;
    }

    protected Int2 readInt2()
    {
        Int2 result = new Int2
        {
            x = readInt(),
            y = readInt()
        };
        return result;
    }

    protected Int3 readInt3()
    {
        Int3 result = new Int3
        {
            x = readInt(),
            y = readInt(),
            z = readInt()
        };
        return result;
    }

    protected Int4 readInt4()
    {
        Int4 result = new Int4
        {
            x = readInt(),
            y = readInt(),
            z = readInt(),
            w = readInt()
        };
        return result;
    }

    protected Int5 readInt5()
    {
        Int5 result = new Int5
        {
            x = readInt(),
            y = readInt(),
            z = readInt(),
            w = readInt(),
            v = readInt()
        };
        return result;
    }

    protected string readString()
    {
        int len = readInt();
        byte[] bytes = readBytes(len);
        string str = "";
        int count = bytes.Length;
        for (int i = 0; i < count; i++)
        {
            str += (char)bytes[i];
        }
        return str;
    }

    protected byte[] readBytes(int len)
    {
        byte[] bytes = new byte[len];
        for (int i = 0; i < len; i++)
        {
            if (m_index >= m_binary.Length)
            {
                bytes[i] = 0;
            }
            else
            {
                bytes[i] = m_binary[m_index];
                m_index++;
            }
        }
        return bytes;
    }
}

public class Int2
{
    public int x = 0;
    public int y = 0;
}

public class Int3 : Int2
{
    public int z = 0;
}

public class Int4 : Int3
{
    public int w = 0;
}

public class Int5 : Int4
{
    public int v = 0;
}

public class Int5 : Int5
{
    public int v = 0;
}