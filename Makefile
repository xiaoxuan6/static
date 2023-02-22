url=http://www.baidu.com
filename=1.png
element=
width=1200
height=800

run:
	 go run main.go exec -u "$(url)" -f "$(filename)" -e "$(element)" -g "$(height)" -w "$(width)"
