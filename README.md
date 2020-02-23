# Configo配置打包工具

将Excel文件打包成对应的文件格式



## 配置文件

### CSharp

    {
        "platforms": [
            {
                "name": "client",
                "ext": "C",
                "output": "binary",
                "outputfolder": true,
                "outputext": "",
                "allinone": false,
                "createcls": true,
                "clsfolder": "",
                "language": "csharp",
                "libcls": "BaseConfigData",
                "libpath": "",
                "datastructs": {
                    "int2": {
                        "libcls": "Int2",
                        "libpath": ""
                    },
                    "int3": {
                        "libcls": "Int3",
                        "libpath": ""
                    },
                    "int4": {
                        "libcls": "Int4",
                        "libpath": ""
                    },
                    "int5": {
                        "libcls": "Int5",
                        "libpath": ""
                    }
                }
            }
        ]
    }


### TypeScript

    {
        "platforms": [
            {
                "name": "client",
                "ext": "C",
                "output": "json",
                "outputfolder": true,
                "outputext": "",
                "allinone": false,
                "createcls": true,
                "clsfolder": "",
                "language": "typescript",
                "libcls": "CCArmor.BaseJsonConfigData",
                "libpath": "",
                "datastructs": {
                    "int2": {
                        "libcls": "CCArmor.Num2",
                        "libpath": ""
                    },
                    "int3": {
                        "libcls": "CCArmor.Num3",
                        "libpath": ""
                    },
                    "int4": {
                        "libcls": "CCArmor.Num4",
                        "libpath": ""
                    },
                    "int5": {
                        "libcls": "CCArmor.Num5",
                        "libpath": ""
                    }
                }
            }
        ]
    }

* platforms: 平台配置

    * name: 平台名称

    * ext: 平台字段后缀，用于过滤特殊平台才需要打包的字段和页签（在对应的字段或标签名称后添加“_ext”）

    * output: 平台生成的文件格式，目前支持json、xml、binary

    * outputfolder: 是否为平台输出文件建立单独的文件夹

    * outputext: 平台输出文件后缀，不填则是用默认后缀（默认后缀分别为：json、xml、bytes）
    
    * allinone: 是否将同一个Excel中的表格打包在同一个文件中（binary不支持）
    
    * createcls: 是否创建配置专用的类文件，目前支持json、binary

    * clsfolder: 配置类文件的目录名称，留空则不创建文件夹

    * language: 生成配置定义文件所使用的语言，目前支持csharp、typescript

    * libcls: 配置文件基类的名称

    * libpath: 配置文件基类的路径

    * datastructs: 特殊数据结构对应的类的名称和路径

## Excel配置

* Excel格式

    * 第一行：文件注释

    * 第二行：字段注释

    * 第三行：字段名称

    * 第四行：字段类型

    * 第五行开始：数据正文

* 默认支持的类型

    * 基础类型：int, string

    * 复合类型（需加入对应数据类文件）：int2, int3, int4, int5

        * 符合类型配置方式：复合类型内部数据使用 竖线（'|'） 分隔

    * 数组类型：任意非数组类型+[]; 如：int[], string[]

        * 数组类型配置方式：数组元素使用 逗号（','）分隔


## 运行

    Configo.exe -c "./config.json" -i "./files.txt"

    * -c 配置文件

    * -i 需要打包的配置文件列表
