package main

import "testing"
import "io/ioutil"
import "fmt"

var lorem2 string = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Massa vitae tortor condimentum lacinia quis vel. Enim tortor at auctor urna nunc id cursus metus aliquam. Platea dictumst vestibulum rhoncus est pellentesque. Nibh ipsum consequat nisl vel. Quam lacus suspendisse faucibus interdum posuere lorem ipsum dolor. Orci a scelerisque purus semper. Nulla pellentesque dignissim enim sit amet venenatis urna. Risus quis varius quam quisque id diam. Eget nullam non nisi est sit amet. Ut faucibus pulvinar elementum integer enim neque. Sit amet consectetur adipiscing elit ut aliquam purus sit amet.

Auctor elit sed vulputate mi sit amet mauris. Id semper risus in hendrerit gravida. Tortor condimentum lacinia quis vel eros donec ac odio. Lacus sed turpis tincidunt id aliquet risus feugiat. Hendrerit dolor magna eget est lorem ipsum dolor sit. Id neque aliquam vestibulum morbi blandit cursus. Id ornare arcu odio ut sem nulla pharetra. Mi in nulla posuere sollicitudin. Ac turpis egestas maecenas pharetra. Quam quisque id diam vel quam elementum pulvinar. Risus in hendrerit gravida rutrum quisque non tellus orci ac. Enim tortor at auctor urna nunc id cursus metus. Euismod nisi porta lorem mollis aliquam ut porttitor leo a. Tellus rutrum tellus pellentesque eu tincidunt tortor aliquam nulla facilisi. Aliquet eget sit amet tellus. Lacus vestibulum sed arcu non odio. Egestas maecenas pharetra convallis posuere morbi leo urna molestie. Pretium lectus quam id leo. Elit scelerisque mauris pellentesque pulvinar pellentesque habitant morbi tristique senectus." +

Elementum curabitur vitae nunc sed velit. Consectetur purus ut faucibus pulvinar elementum integer enim. Facilisis volutpat est velit egestas dui. Amet consectetur adipiscing elit duis tristique sollicitudin nibh. Vitae aliquet nec ullamcorper sit amet. Facilisi nullam vehicula ipsum a. Ac auctor augue mauris augue. Nibh mauris cursus mattis molestie a iaculis at. Ut placerat orci nulla pellentesque dignissim enim sit amet venenatis. Adipiscing enim eu turpis egestas pretium aenean pharetra magna. +

Non nisi est sit amet facilisis magna etiam. Posuere morbi leo urna molestie at elementum. Sed enim ut sem viverra aliquet eget sit amet tellus. Aliquet sagittis id consectetur purus ut. Arcu felis bibendum ut tristique et egestas. Nunc non blandit massa enim nec dui nunc. Cursus metus aliquam eleifend mi in nulla. Elementum sagittis vitae et leo duis ut diam. Enim nunc faucibus a pellentesque. Tellus rutrum tellus pellentesque eu. Placerat in egestas erat imperdiet sed euismod. Nunc vel risus commodo viverra. Nec dui nunc mattis enim ut tellus elementum sagittis vitae. Id leo in vitae turpis massa sed elementum. Nam libero justo laoreet sit. Vitae auctor eu augue ut lectus arcu bibendum at. Eget egestas purus viverra accumsan in nisl nisi. Elementum nisi quis eleifend quam adipiscing vitae proin."

Sed viverra tellus in hac habitasse platea dictumst. Ipsum dolor sit amet consectetur adipiscing. Duis convallis convallis tellus id interdum velit laoreet id donec. Neque egestas congue quisque egestas diam. Non pulvinar neque laoreet suspendisse interdum consectetur libero id. Volutpat sed cras ornare arcu dui vivamus. Aliquam nulla facilisi cras fermentum odio eu. Praesent semper feugiat nibh sed pulvinar proin gravida hendrerit. Quam elementum pulvinar etiam non quam lacus. Tristique senectus et netus et malesuada fames. Tincidunt lobortis feugiat vivamus at augue eget. Duis ut diam quam nulla porttitor massa. Iaculis nunc sed augue lacus viverra vitae congue eu consequat. Laoreet sit amet cursus sit amet dictum sit. Aliquam id diam maecenas ultricies mi.`

func setup() string {
	bytes, err := ioutil.ReadFile("romeo_and_juliet.txt")
	if err != nil {
		fmt.Println("Setup failed", err)
		return ""
	} else {
		fmt.Println("Setup succeeded", err)
		return string(bytes) + string(bytes) + string(bytes) + string(bytes)

	}
}

var lorem string = setup()

func BenchmarkWordCountNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WordCountNaive(lorem)
	}
}

func BenchmarkWordCountAllocEachStringOnlyOnce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WordCountAllocEachStringOnlyOnce(lorem)
	}
}

func BenchmarkWordCountStringsFields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WordCountWithStringsFields(lorem)
	}
}
