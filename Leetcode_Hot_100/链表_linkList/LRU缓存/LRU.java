// 每次get和put都要把对应的结点移动到最前面
// 在get的时候，找不到返回-1，不需要移动；找到先返回值，再移动
// put的时候要注意容积大小，超过容积要删除尾巴，也就是最近不使用的点
// put还要注意是否存在，存在的话要更新里面的结点值，不存在的话需要添加

class LRUCache {
    class DoubleLink{
        // 双向链表
        // prev,next
        public DoubleLink prev;
        public DoubleLink next;
        public int key;
        public int val;
        // 构造函数
        public DoubleLink(){}
        public DoubleLink(int key, int val) {
            this.key = key;
            this.val = val;
        }
    }

    // 初始化好一些变量
    public int capacity;
    public DoubleLink head, tail;
    public HashMap<Integer, DoubleLink> cache = new HashMap<>();
    public int size = 0;
    public LRUCache(int capacity) {
        this.capacity = capacity;
        this.head = new DoubleLink();
        this.tail = new DoubleLink();
        this.head.next = this.tail;
        this.tail.prev = this.head;
    }
    
    public int get(int key) {
        // 先去查询一下存不存在
        DoubleLink ans = cache.get(key);
        if (ans == null) {
            return -1;
        }
        // 移动到头部
        move2Head(ans);
        // 然后返回
        return ans.val;
    }
    
    public void put(int key, int value) {
        // 先看看存不存在
        // 如果存在，那么就更新val并且移动到头部
        DoubleLink ans = cache.get(key);
        if (ans == null) {
            // 说明不存在，添加进去
            ans = new DoubleLink(key, value);
            // 同时也更新哈希表
            cache.put(key,ans);
            size++;
            add2Head(ans);
            if (size > capacity) {
                // 删除末尾
                DoubleLink retail = tail.prev;
                cache.remove(retail.key);
                size--;
                removeTail();
            }
        } else {
            ans.val = value;
	        move2Head(ans);
        }

    }
    private void move2Head(DoubleLink node) {
        // 处理好当前关系
        node.prev.next = node.next;
        node.next.prev = node.prev;
        // 然后移动到头部
        add2Head(node);
    }
    private void add2Head(DoubleLink node) {
        head.next.prev = node;
	    node.next = head.next;
	    node.prev = head;
	    head.next = node;
    }
    private void removeTail() {
        DoubleLink retail = tail.prev;
        retail.prev.next = tail;
        tail.prev = retail.prev;
    }
}