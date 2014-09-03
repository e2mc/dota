混战模型

* 程序环境:
	
		开发环境为 mac + go1.3 
		
* 程序运行方法:

		git clone git@gitlab.mogujie.org:liangxiao/gamberetto-festvial.git
		cd gamberetto-festvial
		export GOPATH=`pwd`
		cd src/main
		go run main.go 
* 说明: 
	
		模型构建了4个结构
		Team     => 队伍  初始化为正义(good)与邪恶(evil)两支队伍 
		hero	 => 英雄
		ob       => 裁判 
		skill    => 技能

		程序最终有三种结果: 正义胜利(good win) 邪恶胜利(evil win) 打平(both win)
	