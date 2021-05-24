using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace MotionPlanner
{
    class JPS
    {
        public class Node : IComparable<Node>
        {
            public int x = 0;
            public int y = 0;
            public double cost = 0;
            public double G = 0; // 移动代价
            public double H = 0; // 

            public List<Motion> forceList = new List<Motion>();

            public Node front = null;
            public Node(int x, int y)
            {
                this.x = x;
                this.y = y;
            }
            public Node(int x, int y, double G)
            {
                this.x = x;
                this.y = y;
                this.G = G;
            }

            public Node Clone()
            {
                Node node = new Node(x, y, G);
                node.cost = cost;
                node.forceList = forceList;
                node.front = front;
                node.H = H;

                return node;
            }

            public int CompareTo(Node other)
            {
                return this.cost < other.cost ? -1 : (this.cost == other.cost ? 0 : 1);
            }

            public bool isEqual(Node other) // 判断输入结点和本结点是否相同
            {
                if (this.x == other.x && this.y == other.y)
                {
                    return true;
                }
                return false;
            }


        }

        PriorityQueue<Node> openlist = new PriorityQueue<Node>();
        GridMap map;
        public JPS(GridMap map)
        {
            this.map = map;
        }
        public void Search()
        {
            Node goal = new Node(map.goal.X, map.goal.Y);
            // 将起点加入openlist
            openlist.Push(new Node(map.origin.X, map.origin.Y));
            openlist.Top().forceList.AddRange(motionList);

            bool flag = false;
            while (openlist.Count != 0)
            {
                if (flag) break;
                Node node = openlist.Pop();

                foreach (Motion m in node.forceList)
                {
                    if (flag) break;
                    Node parent = new Node(node.x, node.y, node.G);
                    parent.cost = node.G + GetEuclideanDistance(node, goal);
                    parent.front = node.front;
                    while (GetNodeStatus(parent) != (int)GridMap.MapStatus.Occupied)
                    {
                        if (flag) break;
                        // 垂直跳跃 /////////////////////////////////////////////////////////////////////
                        Node current = new Node(parent.x, parent.y, parent.G);
                        current.cost = parent.G + GetEuclideanDistance(parent, goal);
                        current.front = parent.front;
                        while (GetNodeStatus(current) != (int)GridMap.MapStatus.Occupied)
                        {
                            if (GetNodeStatus(current) == (int)GridMap.MapStatus.Exploring)
                            {
                                current = new Node(current.x + m.delta_x, current.y, current.G + 1);
                                current.front = parent;
                                current.cost = current.G + GetEuclideanDistance(current, goal);
                                continue;
                            }
                            List<Motion> forceList = GetForcedNeighborList(current, m, 0);
                            if (forceList.Count != 0)
                            {
                                current.forceList.AddRange(forceList);
                                if (!parent.isEqual(node) && !parent.isEqual(current))
                                {
                                    map.map[parent.x][parent.y] = (int)GridMap.MapStatus.Exploring;
                                    //parent.forceList.Add(m);
                                    //openlist.Push(parent);
                                }
                                map.map[current.x][current.y] = (int)GridMap.MapStatus.Exploring;
                                openlist.Push(current);
                                break;
                            }
                            if (GetNodeStatus(current) == (int)GridMap.MapStatus.Unoccupied)
                                map.map[current.x][current.y] = (int)GridMap.MapStatus.Explored;

                            current = new Node(current.x + m.delta_x, current.y, current.G + 1);
                            current.front = parent;
                            current.cost = current.G + GetEuclideanDistance(current, goal);
 
                            if (current.isEqual(goal))
                            {
                                map.map[parent.x][parent.y] = (int)GridMap.MapStatus.Exploring;
                                while (current.front != null)
                                {
                                    //map.map[current.x][current.y] = (int)GridMap.MapStatus.Road;
                                    map.road.Add(new Point(current.x, current.y));
                                    current = current.front;
                                }
                                map.road.Add(map.origin);
                                flag = true;
                                break;
                            }
                            Thread.Sleep(10);
                        }
                        if (flag) break;
                        // 水平跳跃 //////////////////////////////////////////////////////////////////////////////
                        current = new Node(parent.x, parent.y, parent.G);
                        current.cost = parent.G + GetEuclideanDistance(parent, goal);
                        current.front = parent.front;
                        while (GetNodeStatus(current) != (int)GridMap.MapStatus.Occupied)
                        {
                            if (GetNodeStatus(current) == (int)GridMap.MapStatus.Exploring)
                            {
                                current = new Node(current.x, current.y + m.delta_y, current.G + 1);
                                current.front = parent;
                                current.cost = current.G + GetEuclideanDistance(current, goal);
                                continue;
                            }
                            List<Motion> forceList = GetForcedNeighborList(current, m, 1);
                            if (forceList.Count != 0)
                            {
                                current.forceList.AddRange(forceList);
                                if (!parent.isEqual(node) && !parent.isEqual(current))
                                {
                                    map.map[parent.x][parent.y] = (int)GridMap.MapStatus.Exploring;
                                    //parent.forceList.Add(m);
                                    //openlist.Push(parent);
                                }
                                map.map[current.x][current.y] = (int)GridMap.MapStatus.Exploring;
                                openlist.Push(current);
                                break;
                            }
                            if (GetNodeStatus(current) == (int)GridMap.MapStatus.Unoccupied)
                                map.map[current.x][current.y] = (int)GridMap.MapStatus.Explored;

                            current = new Node(current.x, current.y + m.delta_y, current.G + 1);
                            current.front = parent;
                            current.cost = current.G + GetEuclideanDistance(current, goal);

                            if (current.isEqual(goal))
                            {
                                map.map[parent.x][parent.y] = (int)GridMap.MapStatus.Exploring;
                                while (current.front != null)
                                {
                                    //map.map[current.x][current.y] = (int)GridMap.MapStatus.Road;
                                    map.road.Add(new Point(current.x, current.y));
                                    current = current.front;
                                }
                                map.road.Add(map.origin);
                                flag = true;
                                break;
                            }
                            Thread.Sleep(10);
                        }
                        // 对角线跳跃 ////////////////////////////////////////////////////////////////////////////
                        parent = new Node(parent.x + m.delta_x, parent.y + m.delta_y, parent.G + Math.Sqrt(2));
                        parent.front = node;
                        parent.cost = parent.G + GetEuclideanDistance(parent, goal);
                    }
                    
                }
            }
            Thread.Sleep(100);
            map.searchFlag = 1;
        }

        private double GetManhattanDistance(Node p1, Node p2)
        {
            return Math.Abs(p1.x - p2.x) + Math.Abs(p1.y - p2.y);
        }
        private double GetEuclideanDistance(Node p1, Node p2)
        {
            return Math.Sqrt(Math.Pow(p1.x - p2.x, 2) + Math.Pow(p1.y - p2.y, 2));
        }

        private List<Motion> GetForcedNeighborList(Node current, Motion m, int jumpFlag)
        {
            List<Motion> forceList = new List<Motion>();
            if (jumpFlag == 0) // 垂直跳跃
            {
                if (map.map[current.x][current.y + 1] == (int)GridMap.MapStatus.Occupied
                    && map.map[current.x + m.delta_x][current.y + 1] == (int)GridMap.MapStatus.Unoccupied)
                {
                    forceList.Add(new Motion(m.delta_x, 1, Math.Sqrt(2)));
                }
                if (map.map[current.x][current.y - 1] == (int)GridMap.MapStatus.Occupied
                    && map.map[current.x + m.delta_x][current.y - 1] == (int)GridMap.MapStatus.Unoccupied)
                {
                    forceList.Add(new Motion(m.delta_x, -1, Math.Sqrt(2)));
                }

            }
            else if (jumpFlag == 1) // 水平跳跃
            {
                if (map.map[current.x + 1][current.y] == (int)GridMap.MapStatus.Occupied
                    && map.map[current.x + 1][current.y + m.delta_y] == (int)GridMap.MapStatus.Unoccupied)
                {
                    forceList.Add(new Motion(1, m.delta_y, Math.Sqrt(2)));
                }
                if (map.map[current.x - 1][current.y] == (int)GridMap.MapStatus.Occupied
                    && map.map[current.x - 1][current.y + m.delta_y] == (int)GridMap.MapStatus.Unoccupied)
                {
                    forceList.Add(new Motion(-1, m.delta_y, Math.Sqrt(2)));

                }
            }
            return forceList;
        }
        int GetNodeStatus(Node parent)
        {
            return map.map[parent.x][parent.y];
        }

        public class Motion
        {
            public int delta_x;
            public int delta_y;
            public double delta_cost;
            public Motion(int x, int y, double cost)
            {
                delta_x = x;
                delta_y = y;
                delta_cost = cost;
            }
        }

        List<Motion> motionList = new List<Motion>
        {
            new Motion(-1,  1,  Math.Sqrt(2)),
            new Motion( 1,  1,  Math.Sqrt(2)),
            new Motion( 1, -1,  Math.Sqrt(2)),
            new Motion(-1, -1,  Math.Sqrt(2)),
        };


    }
}
