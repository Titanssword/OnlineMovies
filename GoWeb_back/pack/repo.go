package pack

import (
	"fmt"
	"io/ioutil"
)

var currentId int

var movies TodoList

// Give us some seed data
//func init() {
//	RepoCreateTodo(Todo{Id: 11, Name: "阿拉伯的劳伦斯" , Url: "http://10.199.130.73:81/[42].阿拉伯的劳伦斯.Lawrence.of.Arabia.1962.HDTV.MiniSD-TLF.mkv"})
//	RepoCreateTodo(Todo{Id: 12, Name: "异形" ,Url: "http://10.199.130.73:81/[43].异形.Alien.DC.1979.HDTV.MiniSD-TLF.mkv"})
//	RepoCreateTodo(Todo{Id: 13, Name: "天使爱美丽" ,Url: "http://10.199.130.73:81/[44].天使爱美丽.Amelie.2001.BDRip.x264.2AAC.miniSD-TLF.mkv"})
//}

func listFile(folder string){
	files, _ := ioutil.ReadDir(folder) //specify the current dir
	for _,file := range files {
		if file.IsDir() {
			listFile(folder + "/" + file.Name())
		} else {
			movie := Todo{
				Name: file.Name(),
				Url:  "http://10.199.130.73:81/movies/" + file.Name(),
			}
			RepoCreateTodo(movie)
			//fmt.Println(folder + "/" + file.Name())
		}
	}
}

func RepoFindTodo(Id int) Todo {
	for _, t := range movies {
		if t.Id == Id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	movies = append(movies, t)
	return t
}

func RepoDestroyTodo(Id int) error {
	for i, t := range movies {
		if t.Id == Id {
			movies = append(movies[:i], movies[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with Id of %d to delete", Id)
}