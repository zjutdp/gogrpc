// https://qinm08.github.io/2016/160801-spark-mllib-1-naive-bayes/#more

import org.apache.spark.SparkConf
import org.apache.spark.SparkContext
import org.apache.spark.mllib.linalg.Vectors
import org.apache.spark.mllib.regression.LabeledPoint
import org.apache.spark.mllib.classification.NaiveBayes

// Spark 读取文本文件
val rawtxt = sc.textFile("input/human-body-features.txt")
// 将文本文件的内容转化为我们需要的数据结构 LabeledPoint
val allData = rawtxt.map {
    line =>
        val colData = line.split(',')
        LabeledPoint(colData(0).toDouble,
                Vectors.dense(colData(1).split(' ').map(_.toDouble)))
}
// 训练
val nbTrained = NaiveBayes.train(allData)
// 待分类的特征集合
val txt = "6 130 8";
val vec = Vectors.dense(txt.split(' ').map(_.toDouble))
// 预测（分类）
val nbPredict = nbTrained.predict(vec)
println("预测此人性别是：" + (if(nbPredict == 0) "女" else "男"))