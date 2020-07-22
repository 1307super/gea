package main
import(
	"encoding/json"
	"fmt"
)

type Datanode struct {
	Id int `json:"id"`
	ParentId int `json:"pid"`
	Name string `json:"name"`
	Child []*Datanode `json:"child"`
}
var Jsondata []byte  //存储json数据
func main(){
	Data := make([]*Datanode, 0)  //存储所有初始化struct
	a := Datanode{
		Id :0,
		ParentId : -1,
		Name : "目录",
	}
	Data = append(Data,&a)

	b := Datanode{
		Id:1,
		ParentId:0,
		Name :"一、水果",
	}
	Data = append(Data,&b)

	c := Datanode{
		Id:2,
		ParentId:1,
		Name :"1.苹果",
	}
	Data = append(Data,&c)

	d := Datanode{
		Id:3,
		ParentId:1,
		Name :"2.香蕉",
	}
	Data = append(Data,&d)

	e := Datanode{
		Id:4,
		ParentId:0,
		Name :"二、蔬菜",
	}
	Data = append(Data,&e)

	f := Datanode{
		Id:5,
		ParentId:4,
		Name :"1.芹菜",
	}
	Data = append(Data,&f)

	g := Datanode{
		Id:6,
		ParentId:4,
		Name :"2.黄瓜",
	}
	Data = append(Data,&g)

	h := Datanode{
		Id:7,
		ParentId:6,
		Name :"(1)黄瓜特点",
	}
	Data = append(Data,&h)

	i := Datanode{
		Id:8,
		ParentId:4,
		Name :"3.西红柿",
	}
	Data = append(Data,&i)

	j := Datanode{
		Id:9,
		ParentId:0,
		Name :"三、关系",
	}
	Data = append(Data,&j)

	Anode := Data[0] //父节点
	makeTree(Data,Anode)  //调用生成tree
	transformJson(Anode)
	//jsontoTree(Jsondata)
}

func makeTree(Allnode []*Datanode, node *Datanode){        //参数为父节点，添加父节点的子节点指针切片
	childs,_ :=haveChild(Allnode,node)      //判断节点是否有子节点并返回
	if childs != nil{
		fmt.Printf("\n")
		fmt.Println(*node)
		fmt.Println("子节点：")

		for _ ,v:= range childs{
			fmt.Println(*v)
		}                                   //打印

		node.Child = append(node.Child,childs[0:]...)    //添加子节点
		for _,v := range childs{            //查询子节点的子节点，并添加到子节点
			_, has := haveChild(Allnode,v)
			if has {
				makeTree(Allnode, v)          //递归添加节点
			}
		}
	}
}

func haveChild(Allnode []*Datanode, node *Datanode)(childs []*Datanode,yes bool){
	for _,v := range Allnode{
		if v.ParentId == node.Id {
			childs = append(childs, v)
		}
	}
	if childs != nil {
		yes = true
	}
	return
}

func transformJson(Data *Datanode){  //转为json

	Jsondata,_ = json.Marshal(Data)

	fmt.Println(string(Jsondata))
}

func jsontoTree(jsondata []byte){  //json转struct
	var tree_node *Datanode
	err := json.Unmarshal(jsondata,&tree_node)
	fmt.Println("22222222222222",string(jsondata))
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(tree_node.Id,tree_node.ParentId,tree_node.Name)
		printTree(tree_node)

	}
}


func printTree(tree_node *Datanode){
	for _,v := range tree_node.Child {
		fmt.Printf("%d,%d,%s",v.Id,v.ParentId,v.Name)
		fmt.Println("##########")
		if len(v.Child) != 0{
			printTree(v)
		}
	}
}