package main

import (
	"github.com/monkeydioude/moon"
)

func example1(r *moon.Request, c *moon.Configuration) ([]byte, int, error) {
	return []byte("" +
		"                               ..,,,,,,,,,.. \n" +
		"                        .,;%%%%%%%%%%%%%%%%%%%%;,. \n" +
		"                      %%%%%%%%%%%%%%%%%%%%////%%%%%%, .,;%%;, \n" +
		"   		.,;%/,%%%%%/////%%%%%%%%%%%%%%////%%%%,%%//%%%, \n" +
		"   	.,;%%%%/,%%%///%%%%%%%%%%%%%%%%%%%%%%%%%%%%,////%%%%;, \n" +
		"     .,%%%%%%//,%%%%%%%%%%%%%%%%@@%a%%%%%%%%%%%%%%%%,%%/%%%%%%%;, \n" +
		"   .,%//%%%%//,%%%%///////%%%%%%%@@@%%%%%%///////%%%%,%%//%%%%%%%%, \n" +
		" ,%%%%%///%%//,%%//%%%%%///%%%%%@@@%%%%%////%%%%%%%%%,/%%%%%%%%%%%%% \n" +
		".%%%%%%%%%////,%%%%%%%//%///%%%%@@@@%%%////%%/////%%%,/;%%%%%%%%/%%% \n" +
		"%/%%%%%%%/////,%%%%///%%////%%%@@@@@%%%///%%/%%%%%//%,////%%%%//%%%' \n" +
		"%//%%%%%//////,%/%a`  'a%///%%%@@@@@@%%////a`  'a%%%%,//%///%/%%%%% \n" +
		"%///%%%%%%///,%%%%@@aa@@%//%%%@@@@S@@@%%///@@aa@@%%%%%,/%////%%%%% \n" +
		"%%//%%%%%%%//,%%%%%///////%%%@S@@@@SS@@@%%/////%%%%%%%,%////%%%%%' \n" +
		"%%//%%%%%%%//,%%%%/////%%@%@SS@@@@@@@S@@@@%%%%/////%%%,////%%%%%' \n" +
		"`%/%%%%//%%//,%%%///%%%%@@@S@@@@@@@@@@@@@@@S%%%%////%%,///%%%%%' \n" +
		"  %%%%//%%%%/,%%%%%%%%@@@@@@@@@@@@@@@@@@@@@SS@%%%%%%%%,//%%%%%' \n" +
		"  `%%%//%%%%/,%%%%@%@@@@@@@@@@@@@@@@@@@@@@@@@S@@%%%%%,/////%%' \n" +
		"   `%%%//%%%/,%%%@@@SS@@SSs@@@@@@@@@@@@@sSS@@@@@@%%%,//%%//%' \n" +
		"    `%%%%%%/  %%S@@SS@@@@@Ss` .,,.    'sS@@@S@@@@%'  ///%/%' \n" +
		"     `%%%/    %SS@@@@SSS@@S.         .S@@SSS@@@@'    //%%' \n" +
		"              /`S@@@@@@SSSSSs,     ,sSSSSS@@@@@' \n" +
		"   	     %%//`@@@@@@@@@@@@@Ss,sS@@@@@@@@@@@'/ \n" +
		"   	   %%%%@@00`@@@@@@@@@@@@@'@@@@@@@@@@@'//%% \n" +
		"      %%%%%%a%@@@@000aaaaaaaaa00a00aaaaaaa00%@%%%%% \n" +
		"   %%%%%%a%%@@@@@@@@@@000000000000000000@@@%@@%%%@%%% \n" +
		" %%%%%%a%%@@@%@@@@@@@@@@@00000000000000@@@@@@@@@%@@%%@%% \n" +
		"%%%aa%@@@@@@@@@@@@@@0000000000000000000000@@@@@@@@%@@@%%%% \n" +
		"%%@@@@@@@@@@@@@@@00000000000000000000000000000@@@@@@@@@%%%%%"), 200, nil
}

func main() {

	h := moon.Moon()
	// Me API es su API
	h.WithHeader("Access-Control-Allow-Origin", "*")

	// Will call example1() func every time a GET on "/example1" URI is caught
	h.Routes.AddGet("example1", example1)
	moon.ServerRun(":8080", h)
}
