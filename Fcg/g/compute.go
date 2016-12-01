package g

//"fmt"

func CompRout(net string, node string) {
	AddrMap := make(map[string]Nodelinkmap)
	AddrMap = NodeMap[net] //获取子网内全部机器信息
	srcnodecurrent := AddrMap[node]

	srcPublicTele := IPinfo{}
	srcPublicTele.Ip = ""
	srcPublicTele.Isp = ""
	srcPublicTele.Status = 1

	srcPublicMobile := IPinfo{}
	srcPublicMobile.Ip = ""
	srcPublicMobile.Isp = ""
	srcPublicMobile.Status = 1

	srcPublicUni := IPinfo{}
	srcPublicUni.Ip = ""
	srcPublicUni.Isp = ""
	srcPublicUni.Status = 1

	srcPrivate := IPinfo{}
	srcPrivate.Ip = ""
	srcPrivate.Isp = ""
	srcPrivate.Status = 1

	srcPublic := IPinfo{}
	srcPublic.Ip = ""
	srcPublic.Isp = ""
	srcPublic.Status = 1

	for _, ip := range srcnodecurrent.Ippublic { //找到能用的 公网ip
		if ip.Status == 1 {
			srcPublic = ip
			break
		}
	}

	for _, ip := range srcnodecurrent.Ipprivate { //找到能用的 私有ip
		if ip.Status == 1 {
			srcPrivate = ip
			break
		}
	}

	for _, ip := range srcnodecurrent.Ippublic {
		if (ip.Isp == "电信") && (ip.Status == 1) {
			srcPublicTele = ip
		}
		if (ip.Isp == "移动") && (ip.Status == 1) {
			srcPublicMobile = ip
		}
		if (ip.Isp == "联通") && (ip.Status == 1) {
			srcPublicUni = ip
		}
	}

	if srcnodecurrent.Status == 0 {
		if _, ok := RouteMap[net][node]; ok { //如果原路径中有该路 判断是否是人为改动
			delete(RouteMap[net], node)
		}

	} else {
		if net == "1" { //如果为国内地区
			if srcnodecurrent.HostType == "vm" {
				for nodetmp, nodeinfo := range AddrMap {
					if node != nodetmp {
						if NodeMap[net][nodetmp].Status == 0 {
							if _, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
								//RouteMap[net][node]
								delete(RouteMap[net][node], nodeinfo.Id)
							}
						} else {
							if nodeinfo.HostType == "vm" { //如果dest为vm
								routnodelink := RoutlinkMap{}
								//fmt.Println("nodeinfo:  ", nodeinfo)
								dstPublicTele := IPinfo{}
								dstPublicTele.Ip = ""
								dstPublicTele.Isp = ""
								dstPublicTele.Status = 1

								dstPublicMobile := IPinfo{}
								dstPublicMobile.Ip = ""
								dstPublicMobile.Isp = ""
								dstPublicMobile.Status = 1

								dstPublicUni := IPinfo{}
								dstPublicUni.Ip = ""
								dstPublicUni.Isp = ""
								dstPublicUni.Status = 1

								for _, ip := range nodeinfo.Ippublic {
									if (ip.Isp == "电信") && (ip.Status == 1) {
										dstPublicTele = ip
									}
									if (ip.Isp == "移动") && (ip.Status == 1) {
										dstPublicMobile = ip
									}
									if (ip.Isp == "联通") && (ip.Status == 1) {
										dstPublicUni = ip
									}
								}

								if (srcPublicTele.Ip != "") && (dstPublicTele.Ip != "") {
									routnodelink.From.Ip = srcPublicTele.Ip
									routnodelink.From.Isp = srcPublicTele.Isp
									routnodelink.From.Location = srcnodecurrent.Location
									routnodelink.From.NodeId = srcnodecurrent.Id

									routnodelink.Level = 1
									routnodelink.Lost = ""
									routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
									routnodelink.Delays = 300

									routnodelink.To.Ip = dstPublicTele.Ip
									routnodelink.To.Isp = dstPublicTele.Isp
									routnodelink.To.Location = nodeinfo.Location
									routnodelink.To.NodeId = nodeinfo.Id

									//fmt.Println("ip", routnodelink.To.Ip)
								} else {
									if (srcPublicMobile.Ip != "") && (dstPublicMobile.Ip != "") {
										routnodelink.From.Ip = srcPublicMobile.Ip
										routnodelink.From.Isp = srcPublicMobile.Isp
										routnodelink.From.Location = srcnodecurrent.Location
										routnodelink.From.NodeId = srcnodecurrent.Id

										routnodelink.Level = 1
										routnodelink.Lost = ""
										routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
										routnodelink.Delays = 300

										routnodelink.To.Ip = dstPublicMobile.Ip
										routnodelink.To.Isp = dstPublicMobile.Isp
										routnodelink.To.Location = nodeinfo.Location
										routnodelink.To.NodeId = nodeinfo.Id
									} else {
										if (srcPublicUni.Ip != "") && (dstPublicUni.Ip != "") {
											routnodelink.From.Ip = srcPublicUni.Ip
											routnodelink.From.Isp = srcPublicUni.Isp
											routnodelink.From.Location = srcnodecurrent.Location
											routnodelink.From.NodeId = srcnodecurrent.Id

											routnodelink.Level = 1
											routnodelink.Lost = ""
											routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
											routnodelink.Delays = 300

											routnodelink.To.Ip = dstPublicUni.Ip
											routnodelink.To.Isp = dstPublicUni.Isp
											routnodelink.To.Location = nodeinfo.Location
											routnodelink.To.NodeId = nodeinfo.Id
										}
									}
								}
								if v, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
									if v.Status == 1 { //如果是人为改动 则忽略继续
										continue
									} else {
										RouteMap[net][node][nodeinfo.Id] = routnodelink
									}
								} else {
									if _, ok := RouteMap[net]; ok {
										if _, ok := RouteMap[net][node]; ok {
											RouteMap[net][node][nodeinfo.Id] = routnodelink

										} else {
											dstmap := make(map[string]RoutlinkMap)
											dstmap[nodeinfo.Id] = routnodelink
											RouteMap[net][node] = dstmap
										}
									} else {
										dstmap := make(map[string]RoutlinkMap)
										dstmap[nodeinfo.Id] = routnodelink
										nodemap := make(map[string]map[string]RoutlinkMap)
										nodemap[node] = dstmap
										RouteMap[net] = nodemap
										//fmt.Println("hello", nodemap)
										//	生成 cfg 下发
									}
								}
								//
								//fmt.Println(RouteMap)

							}
							if (nodeinfo.HostType == "al0") || (nodeinfo.HostType == "al1") { //和阿里0机器连接 发送到外网ip
								routnodelink := RoutlinkMap{}
								if (srcPublicTele.Ip != "") && (nodeinfo.Ip != "") {
									routnodelink.From.Ip = srcPublicTele.Ip
									routnodelink.From.Isp = srcPublicTele.Isp
									routnodelink.From.Location = srcnodecurrent.Location
									routnodelink.From.NodeId = srcnodecurrent.Id

									routnodelink.Level = 1
									routnodelink.Lost = ""
									routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
									routnodelink.Delays = 300

									routnodelink.To.Ip = nodeinfo.Ip
									routnodelink.To.Isp = ""
									routnodelink.To.Location = nodeinfo.Location
									routnodelink.To.NodeId = nodeinfo.Id
								} else {
									if (srcPublicMobile.Ip != "") && (nodeinfo.Ip != "") {
										routnodelink.From.Ip = srcPublicMobile.Ip
										routnodelink.From.Isp = srcPublicMobile.Isp
										routnodelink.From.Location = srcnodecurrent.Location
										routnodelink.From.NodeId = srcnodecurrent.Id

										routnodelink.Level = 1
										routnodelink.Lost = ""
										routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
										routnodelink.Delays = 300

										routnodelink.To.Ip = nodeinfo.Ip
										routnodelink.To.Isp = ""
										routnodelink.To.Location = nodeinfo.Location
										routnodelink.To.NodeId = nodeinfo.Id
									} else {
										if (srcPublicUni.Ip != "") && (nodeinfo.Ip != "") {
											routnodelink.From.Ip = srcPublicUni.Ip
											routnodelink.From.Isp = srcPublicUni.Isp
											routnodelink.From.Location = srcnodecurrent.Location
											routnodelink.From.NodeId = srcnodecurrent.Id

											routnodelink.Level = 1
											routnodelink.Lost = ""
											routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
											routnodelink.Delays = 300

											routnodelink.To.Ip = nodeinfo.Ip
											routnodelink.To.Isp = ""
											routnodelink.To.Location = nodeinfo.Location
											routnodelink.To.NodeId = nodeinfo.Id
										}
									}
								}
								if v, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
									if v.Status == 1 { //如果是人为改动 则忽略继续
										continue
									} else {
										RouteMap[net][node][nodeinfo.Id] = routnodelink
									}
								} else {
									if _, ok := RouteMap[net]; ok {
										if _, ok := RouteMap[net][node]; ok {
											RouteMap[net][node][nodeinfo.Id] = routnodelink

										} else {
											dstmap := make(map[string]RoutlinkMap)
											dstmap[nodeinfo.Id] = routnodelink
											RouteMap[net][node] = dstmap
										}
									} else {
										dstmap := make(map[string]RoutlinkMap)
										dstmap[nodeinfo.Id] = routnodelink
										nodemap := make(map[string]map[string]RoutlinkMap)
										nodemap[node] = dstmap
										RouteMap[net] = nodemap
										//	fmt.Println("hello", nodemap)
										//	生成 cfg 下发
									}
								}
							}
						}
					}
				}
			}
			if srcnodecurrent.HostType == "al0" {
				for nodetmp, nodeinfo := range AddrMap {
					if node != nodetmp {
						if NodeMap[net][nodetmp].Status == 0 {
							if _, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
								delete(RouteMap[net][node], nodeinfo.Id)
							}
						} else {

							if nodeinfo.HostType == "vm" {
								routnodelink := RoutlinkMap{}

								dstPublicTele := IPinfo{}
								dstPublicTele.Ip = ""
								dstPublicTele.Isp = ""
								dstPublicTele.Status = 1

								dstPublicMobile := IPinfo{}
								dstPublicMobile.Ip = ""
								dstPublicMobile.Isp = ""
								dstPublicMobile.Status = 1

								dstPublicUni := IPinfo{}
								dstPublicUni.Ip = ""
								dstPublicUni.Isp = ""
								dstPublicUni.Status = 1

								for _, ip := range nodeinfo.Ippublic {
									if (ip.Isp == "电信") && (ip.Status == 1) {
										dstPublicTele = ip
									}
									if (ip.Isp == "移动") && (ip.Status == 1) {
										dstPublicMobile = ip
									}
									if (ip.Isp == "联通") && (ip.Status == 1) {
										dstPublicUni = ip
									}
								}
								if (srcPrivate.Ip != "") && (dstPublicTele.Ip != "") {
									routnodelink.From.Ip = srcPrivate.Ip
									routnodelink.From.Isp = ""
									routnodelink.From.Location = srcnodecurrent.Location
									routnodelink.From.NodeId = srcnodecurrent.Id

									routnodelink.Level = 1
									routnodelink.Lost = ""
									routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
									routnodelink.Delays = 300

									routnodelink.To.Ip = dstPublicTele.Ip
									routnodelink.To.Isp = dstPublicTele.Isp
									routnodelink.To.Location = nodeinfo.Location
									routnodelink.To.NodeId = nodeinfo.Id
								} else {
									if (srcPrivate.Ip != "") && (dstPublicMobile.Ip != "") {
										routnodelink.From.Ip = srcPrivate.Ip
										routnodelink.From.Isp = ""
										routnodelink.From.Location = srcnodecurrent.Location
										routnodelink.From.NodeId = srcnodecurrent.Id

										routnodelink.Level = 1
										routnodelink.Lost = ""
										routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
										routnodelink.Delays = 300

										routnodelink.To.Ip = dstPublicMobile.Ip
										routnodelink.To.Isp = dstPublicMobile.Isp
										routnodelink.To.Location = nodeinfo.Location
										routnodelink.To.NodeId = nodeinfo.Id
									} else {
										if (srcPrivate.Ip != "") && (dstPublicMobile.Ip != "") {
											routnodelink.From.Ip = srcPrivate.Ip
											routnodelink.From.Isp = ""
											routnodelink.From.Location = srcnodecurrent.Location
											routnodelink.From.NodeId = srcnodecurrent.Id

											routnodelink.Level = 1
											routnodelink.Lost = ""
											routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
											routnodelink.Delays = 300

											routnodelink.To.Ip = dstPublicUni.Ip
											routnodelink.To.Isp = dstPublicUni.Isp
											routnodelink.To.Location = nodeinfo.Location
											routnodelink.To.NodeId = nodeinfo.Id
										}
									}
								}

								//						if v, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
								//							if v.Status == 1 { //如果是人为改动 则忽略继续
								//								continue
								//							} else {
								//								RouteMap[net][node][nodeinfo.Id] = routnodelink
								//							}
								//						} else {
								//							//RouteMap = make(map[string]map[string]map[string]RoutlinkMap)

								//							if dstmap, ok := RouteMap[net][node]; ok {
								//								dstmap[nodeinfo.Id] = routnodelink
								//								RouteMap[net][node] = dstmap
								//							} else {
								//								if nodemap, ok := RouteMap[net]; ok {
								//									dstmap := make(map[string]RoutlinkMap)
								//									dstmap[nodeinfo.Id] = routnodelink
								//									nodemap[node] = dstmap
								//									RouteMap[net] = nodemap
								//								} else {
								//									dstmap := make(map[string]RoutlinkMap)
								//									dstmap[nodeinfo.Id] = routnodelink
								//									nodemap := make(map[string]map[string]RoutlinkMap)
								//									nodemap[node] = dstmap
								//									RouteMap[net] = nodemap
								//								}
								//							}
								//						}
								if v, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
									if v.Status == 1 { //如果是人为改动 则忽略继续
										continue
									} else {
										RouteMap[net][node][nodeinfo.Id] = routnodelink
									}
								} else {
									if _, ok := RouteMap[net]; ok {
										if _, ok := RouteMap[net][node]; ok {
											RouteMap[net][node][nodeinfo.Id] = routnodelink

										} else {
											dstmap := make(map[string]RoutlinkMap)
											dstmap[nodeinfo.Id] = routnodelink
											RouteMap[net][node] = dstmap
										}
									} else {
										dstmap := make(map[string]RoutlinkMap)
										dstmap[nodeinfo.Id] = routnodelink
										nodemap := make(map[string]map[string]RoutlinkMap)
										nodemap[node] = dstmap
										RouteMap[net] = nodemap
										//	fmt.Println("hello", nodemap)
										//	生成 cfg 下发
									}
								}

							}
							if (nodeinfo.HostType == "al0") || (nodeinfo.HostType == "al1") {
								routnodelink := RoutlinkMap{}
								//找到能用的ip
								dstPublic := IPinfo{}
								dstPublic.Ip = ""
								dstPublic.Isp = ""
								dstPublic.Status = 1

								for _, ip := range nodeinfo.Ippublic { //找到能用的 gongwang ip
									if ip.Status == 1 {
										dstPublic = ip
										break
									}
								}

								if (srcPrivate.Ip != "") && (dstPublic.Ip != "") {
									routnodelink.From.Ip = srcPrivate.Ip
									routnodelink.From.Isp = ""
									routnodelink.From.Location = srcnodecurrent.Location
									routnodelink.From.NodeId = srcnodecurrent.Id

									routnodelink.Level = 1
									routnodelink.Lost = ""
									routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
									routnodelink.Delays = 300

									routnodelink.To.Ip = dstPublic.Ip
									routnodelink.To.Isp = dstPublic.Isp
									routnodelink.To.Location = nodeinfo.Location
									routnodelink.To.NodeId = nodeinfo.Id
								}

								if v, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
									if v.Status == 1 { //如果是人为改动 则忽略继续
										continue
									} else {
										RouteMap[net][node][nodeinfo.Id] = routnodelink
									}
								} else {
									if _, ok := RouteMap[net]; ok {
										if _, ok := RouteMap[net][node]; ok {
											RouteMap[net][node][nodeinfo.Id] = routnodelink

										} else {
											dstmap := make(map[string]RoutlinkMap)
											dstmap[nodeinfo.Id] = routnodelink
											RouteMap[net][node] = dstmap
										}
									} else {
										dstmap := make(map[string]RoutlinkMap)
										dstmap[nodeinfo.Id] = routnodelink
										nodemap := make(map[string]map[string]RoutlinkMap)
										nodemap[node] = dstmap
										RouteMap[net] = nodemap
										//fmt.Println("hello", nodemap)
										//	生成 cfg 下发
									}
								}

							}
						}
					}
				}
			}

			if srcnodecurrent.HostType == "al1" {
				for nodetmp, nodeinfo := range AddrMap {
					if node != nodetmp {
						if NodeMap[net][nodetmp].Status == 0 {
							if _, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
								delete(RouteMap[net][node], nodeinfo.Id)
							}
						} else {
							if nodeinfo.HostType != "vm" {
								routnodelink := RoutlinkMap{}
								dstPublic := IPinfo{}
								dstPublic.Ip = ""
								dstPublic.Isp = ""
								dstPublic.Status = 1

								for _, ip := range nodeinfo.Ippublic { //找到能用的 g公网ip
									if ip.Status == 1 {
										dstPublic = ip
										break
									}
								}

								if (srcPublic.Ip != "") && (dstPublic.Ip != "") {
									routnodelink.From.Ip = srcPublic.Ip
									routnodelink.From.Isp = srcPublic.Isp
									routnodelink.From.Location = srcnodecurrent.Location
									routnodelink.From.NodeId = srcnodecurrent.Id

									routnodelink.Level = 1
									routnodelink.Lost = ""
									routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
									routnodelink.Delays = 300

									routnodelink.To.Ip = dstPublic.Ip
									routnodelink.To.Isp = dstPublic.Isp
									routnodelink.To.Location = nodeinfo.Location
									routnodelink.To.NodeId = nodeinfo.Id
								}
								if v, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
									if v.Status == 1 { //如果是人为改动 则忽略继续
										continue
									} else {
										RouteMap[net][node][nodeinfo.Id] = routnodelink
									}
								} else {
									//RouteMap = make(map[string]map[string]map[string]RoutlinkMap)

									if dstmap, ok := RouteMap[net][node]; ok {
										dstmap[nodeinfo.Id] = routnodelink
										RouteMap[net][node] = dstmap
									} else {
										if nodemap, ok := RouteMap[net]; ok {
											dstmap := make(map[string]RoutlinkMap)
											dstmap[nodeinfo.Id] = routnodelink
											nodemap[node] = dstmap
											RouteMap[net] = nodemap
										} else {
											dstmap := make(map[string]RoutlinkMap)
											dstmap[nodeinfo.Id] = routnodelink
											nodemap := make(map[string]map[string]RoutlinkMap)
											nodemap[node] = dstmap
											RouteMap[net] = nodemap
										}
									}
								}
							} else {
								routnodelink := RoutlinkMap{}

								dstPublicTele := IPinfo{}
								dstPublicTele.Ip = ""
								dstPublicTele.Isp = ""
								dstPublicTele.Status = 1

								dstPublicMobile := IPinfo{}
								dstPublicMobile.Ip = ""
								dstPublicMobile.Isp = ""
								dstPublicMobile.Status = 1

								dstPublicUni := IPinfo{}
								dstPublicUni.Ip = ""
								dstPublicUni.Isp = ""
								dstPublicUni.Status = 1

								for _, ip := range nodeinfo.Ippublic {
									if (ip.Isp == "电信") && (ip.Status == 1) {
										dstPublicTele = ip
									}
									if (ip.Isp == "移动") && (ip.Status == 1) {
										dstPublicMobile = ip
									}
									if (ip.Isp == "联通") && (ip.Status == 1) {
										dstPublicUni = ip
									}
								}
								if (srcPublic.Ip != "") && (dstPublicTele.Ip != "") {
									routnodelink.From.Ip = srcPublic.Ip
									routnodelink.From.Isp = srcPublic.Isp
									routnodelink.From.Location = srcnodecurrent.Location
									routnodelink.From.NodeId = srcnodecurrent.Id

									routnodelink.Level = 1
									routnodelink.Lost = ""
									routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
									routnodelink.Delays = 300

									routnodelink.To.Ip = dstPublicTele.Ip
									routnodelink.To.Isp = dstPublicTele.Isp
									routnodelink.To.Location = nodeinfo.Location
									routnodelink.To.NodeId = nodeinfo.Id
								} else {
									if (srcPublic.Ip != "") && (dstPublicMobile.Ip != "") {
										routnodelink.From.Ip = srcPublic.Ip
										routnodelink.From.Isp = srcPublic.Isp
										routnodelink.From.Location = srcnodecurrent.Location
										routnodelink.From.NodeId = srcnodecurrent.Id

										routnodelink.Level = 1
										routnodelink.Lost = ""
										routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
										routnodelink.Delays = 300

										routnodelink.To.Ip = dstPublicMobile.Ip
										routnodelink.To.Isp = dstPublicMobile.Isp
										routnodelink.To.Location = nodeinfo.Location
										routnodelink.To.NodeId = nodeinfo.Id
									} else {
										if (srcPublic.Ip != "") && (dstPublicMobile.Ip != "") {
											routnodelink.From.Ip = srcPublic.Ip
											routnodelink.From.Isp = srcPublic.Isp
											routnodelink.From.Location = srcnodecurrent.Location
											routnodelink.From.NodeId = srcnodecurrent.Id

											routnodelink.Level = 1
											routnodelink.Lost = ""
											routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
											routnodelink.Delays = 300

											routnodelink.To.Ip = dstPublicUni.Ip
											routnodelink.To.Isp = dstPublicUni.Isp
											routnodelink.To.Location = nodeinfo.Location
											routnodelink.To.NodeId = nodeinfo.Id
										}
									}
								}

								//						if v, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
								//							if v.Status == 1 { //如果是人为改动 则忽略继续
								//								continue
								//							} else {
								//								RouteMap[net][node][nodeinfo.Id] = routnodelink
								//							}
								//						} else {
								//							//RouteMap = make(map[string]map[string]map[string]RoutlinkMap)

								//							if dstmap, ok := RouteMap[net][node]; ok {
								//								dstmap[nodeinfo.Id] = routnodelink
								//								RouteMap[net][node] = dstmap
								//							} else {
								//								if nodemap, ok := RouteMap[net]; ok {
								//									dstmap := make(map[string]RoutlinkMap)
								//									dstmap[nodeinfo.Id] = routnodelink
								//									nodemap[node] = dstmap
								//									RouteMap[net] = nodemap
								//								} else {
								//									dstmap := make(map[string]RoutlinkMap)
								//									dstmap[nodeinfo.Id] = routnodelink
								//									nodemap := make(map[string]map[string]RoutlinkMap)
								//									nodemap[node] = dstmap
								//									RouteMap[net] = nodemap
								//								}
								//							}
								//						}
								if v, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
									if v.Status == 1 { //如果是人为改动 则忽略继续
										continue
									} else {
										RouteMap[net][node][nodeinfo.Id] = routnodelink
									}
								} else {
									if _, ok := RouteMap[net]; ok {
										if _, ok := RouteMap[net][node]; ok {
											RouteMap[net][node][nodeinfo.Id] = routnodelink

										} else {
											dstmap := make(map[string]RoutlinkMap)
											dstmap[nodeinfo.Id] = routnodelink
											RouteMap[net][node] = dstmap
										}
									} else {
										dstmap := make(map[string]RoutlinkMap)
										dstmap[nodeinfo.Id] = routnodelink
										nodemap := make(map[string]map[string]RoutlinkMap)
										nodemap[node] = dstmap
										RouteMap[net] = nodemap
										//	fmt.Println("hello", nodemap)
										//	生成 cfg 下发
									}
								}
							}

						}
					}
				}
			}
		} else { //为国外地区
			if srcnodecurrent.HostType == "zx" {

				for nodetmp, nodeinfo := range AddrMap {
					if node != nodetmp {
						if NodeMap[net][nodetmp].Status == 0 {
							if _, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
								delete(RouteMap[net][node], nodeinfo.Id)
							}
						} else {
							routnodelink := RoutlinkMap{}
							dstPublic := IPinfo{}
							dstPublic.Ip = ""
							dstPublic.Isp = ""
							dstPublic.Status = 1

							for _, ip := range nodeinfo.Ippublic { //找到能用的 g公网ip
								if ip.Status == 1 {
									dstPublic = ip
									break
								}
							}

							if (srcPublic.Ip != "") && (dstPublic.Ip != "") {
								routnodelink.From.Ip = srcPublic.Ip
								routnodelink.From.Isp = srcPublic.Isp
								routnodelink.From.Location = srcnodecurrent.Location
								routnodelink.From.NodeId = srcnodecurrent.Id

								routnodelink.Level = 1
								routnodelink.Lost = ""
								routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
								routnodelink.Delays = 300

								routnodelink.To.Ip = dstPublic.Ip
								routnodelink.To.Isp = dstPublic.Isp
								routnodelink.To.Location = nodeinfo.Location
								routnodelink.To.NodeId = nodeinfo.Id
							}
							//					if v, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
							//						if v.Status == 1 { //如果是人为改动 则忽略继续
							//							continue
							//						} else {
							//							RouteMap[net][node][nodeinfo.Id] = routnodelink
							//						}
							//					} else {
							//						if _, ok := RouteMap[net]; ok {
							//							if _, ok := RouteMap[net][node]; ok {
							//								RouteMap[net][node][nodeinfo.Id] = routnodelink

							//							} else {
							//								dstmap := make(map[string]RoutlinkMap)
							//								dstmap[nodeinfo.Id] = routnodelink
							//								RouteMap[net][node] = dstmap
							//							}
							//						} else {
							//							dstmap := make(map[string]RoutlinkMap)
							//							dstmap[nodeinfo.Id] = routnodelink
							//							nodemap := make(map[string]map[string]RoutlinkMap)
							//							nodemap[node] = dstmap
							//							RouteMap[net] = nodemap
							//							//	生成 cfg 下发
							//						}
							//					}
							if v, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
								if v.Status == 1 { //如果是人为改动 则忽略继续
									continue
								} else {
									RouteMap[net][node][nodeinfo.Id] = routnodelink
								}
							} else {
								if _, ok := RouteMap[net]; ok {
									if _, ok := RouteMap[net][node]; ok {
										RouteMap[net][node][nodeinfo.Id] = routnodelink

									} else {
										dstmap := make(map[string]RoutlinkMap)
										dstmap[nodeinfo.Id] = routnodelink
										RouteMap[net][node] = dstmap
									}
								} else {
									dstmap := make(map[string]RoutlinkMap)
									dstmap[nodeinfo.Id] = routnodelink
									nodemap := make(map[string]map[string]RoutlinkMap)
									nodemap[node] = dstmap
									RouteMap[net] = nodemap
									//	fmt.Println("hello", nodemap)
									//	生成 cfg 下发
								}
							}

						}
					}
				}
			}

			if srcnodecurrent.HostType == "sl" {

				for nodetmp, nodeinfo := range AddrMap {
					if node != nodetmp {

						if NodeMap[net][nodetmp].Status == 0 {
							if _, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
								delete(RouteMap[net][node], nodeinfo.Id)
							}
						} else {
							if nodeinfo.HostType != "zx" {
								routnodelink := RoutlinkMap{}
								dstPrivate := IPinfo{}
								dstPrivate.Ip = ""
								dstPrivate.Isp = ""
								dstPrivate.Status = 1

								for _, ip := range nodeinfo.Ipprivate { //找到能用的 g公网ip
									if ip.Status == 1 {
										dstPrivate = ip
										break
									}
								}

								if (srcPrivate.Ip != "") && (dstPrivate.Ip != "") {
									routnodelink.From.Ip = srcPrivate.Ip
									routnodelink.From.Isp = srcPrivate.Isp
									routnodelink.From.Location = srcnodecurrent.Location
									routnodelink.From.NodeId = srcnodecurrent.Id

									routnodelink.Level = 1
									routnodelink.Lost = ""
									routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
									routnodelink.Delays = 300

									routnodelink.To.Ip = dstPrivate.Ip
									routnodelink.To.Isp = dstPrivate.Isp
									routnodelink.To.Location = nodeinfo.Location
									routnodelink.To.NodeId = nodeinfo.Id
								}
								//						if v, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
								//							if v.Status == 1 { //如果是人为改动 则忽略继续
								//								continue
								//							} else {
								//								RouteMap[net][node][nodeinfo.Id] = routnodelink
								//							}
								//						} else {
								//							//RouteMap = make(map[string]map[string]map[string]RoutlinkMap)

								//							if dstmap, ok := RouteMap[net][node]; ok {
								//								dstmap[nodeinfo.Id] = routnodelink
								//								RouteMap[net][node] = dstmap
								//							} else {
								//								if nodemap, ok := RouteMap[net]; ok {
								//									dstmap := make(map[string]RoutlinkMap)
								//									dstmap[nodeinfo.Id] = routnodelink
								//									nodemap[node] = dstmap
								//									RouteMap[net] = nodemap
								//								} else {
								//									dstmap := make(map[string]RoutlinkMap)
								//									dstmap[nodeinfo.Id] = routnodelink
								//									nodemap := make(map[string]map[string]RoutlinkMap)
								//									nodemap[node] = dstmap
								//									RouteMap[net] = nodemap
								//								}
								//							}
								//						}
								if v, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
									if v.Status == 1 { //如果是人为改动 则忽略继续
										continue
									} else {
										RouteMap[net][node][nodeinfo.Id] = routnodelink
									}
								} else {
									if _, ok := RouteMap[net]; ok {
										if _, ok := RouteMap[net][node]; ok {
											RouteMap[net][node][nodeinfo.Id] = routnodelink

										} else {
											dstmap := make(map[string]RoutlinkMap)
											dstmap[nodeinfo.Id] = routnodelink
											RouteMap[net][node] = dstmap
										}
									} else {
										dstmap := make(map[string]RoutlinkMap)
										dstmap[nodeinfo.Id] = routnodelink
										nodemap := make(map[string]map[string]RoutlinkMap)
										nodemap[node] = dstmap
										RouteMap[net] = nodemap
										//fmt.Println("hello", nodemap)
										//	生成 cfg 下发
									}
								}
							} else {
								routnodelink := RoutlinkMap{}
								dstPublic := IPinfo{}
								dstPublic.Ip = ""
								dstPublic.Isp = ""
								dstPublic.Status = 1

								for _, ip := range nodeinfo.Ippublic { //找到能用的 g公网ip
									if ip.Status == 1 {
										dstPublic = ip
										break
									}
								}

								if (srcPublic.Ip != "") && (dstPublic.Ip != "") {
									routnodelink.From.Ip = srcPublic.Ip
									routnodelink.From.Isp = srcPublic.Isp
									routnodelink.From.Location = srcnodecurrent.Location
									routnodelink.From.NodeId = srcnodecurrent.Id

									routnodelink.Level = 1
									routnodelink.Lost = ""
									routnodelink.Status = 0 //默认为0表示自动生成的  1 为认为修
									routnodelink.Delays = 300

									routnodelink.To.Ip = dstPublic.Ip
									routnodelink.To.Isp = dstPublic.Isp
									routnodelink.To.Location = nodeinfo.Location
									routnodelink.To.NodeId = nodeinfo.Id
								}
								//						if v, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
								//							if v.Status == 1 { //如果是人为改动 则忽略继续
								//								continue
								//							} else {
								//								RouteMap[net][node][nodeinfo.Id] = routnodelink
								//							}
								//						} else {
								//							//RouteMap = make(map[string]map[string]map[string]RoutlinkMap)

								//							if dstmap, ok := RouteMap[net][node]; ok {
								//								dstmap[nodeinfo.Id] = routnodelink
								//								RouteMap[net][node] = dstmap
								//							} else {
								//								if nodemap, ok := RouteMap[net]; ok {
								//									dstmap := make(map[string]RoutlinkMap)
								//									dstmap[nodeinfo.Id] = routnodelink
								//									nodemap[node] = dstmap
								//									RouteMap[net] = nodemap
								//								} else {
								//									dstmap := make(map[string]RoutlinkMap)
								//									dstmap[nodeinfo.Id] = routnodelink
								//									nodemap := make(map[string]map[string]RoutlinkMap)
								//									nodemap[node] = dstmap
								//									RouteMap[net] = nodemap
								//								}
								//							}
								//						}
								if v, ok := RouteMap[net][node][nodeinfo.Id]; ok { //如果原路径中有该路 判断是否是人为改动
									if v.Status == 1 { //如果是人为改动 则忽略继续
										continue
									} else {
										RouteMap[net][node][nodeinfo.Id] = routnodelink
									}
								} else {
									if _, ok := RouteMap[net]; ok {
										if _, ok := RouteMap[net][node]; ok {
											RouteMap[net][node][nodeinfo.Id] = routnodelink

										} else {
											dstmap := make(map[string]RoutlinkMap)
											dstmap[nodeinfo.Id] = routnodelink
											RouteMap[net][node] = dstmap
										}
									} else {
										dstmap := make(map[string]RoutlinkMap)
										dstmap[nodeinfo.Id] = routnodelink
										nodemap := make(map[string]map[string]RoutlinkMap)
										nodemap[node] = dstmap
										RouteMap[net] = nodemap
										//	fmt.Println("hello", nodemap)
										//	生成 cfg 下发
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
