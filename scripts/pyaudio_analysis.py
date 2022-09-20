from pyAudioAnalysis.audioTrainTest import extract_features_and_train
mt, st = 1.0, 0.05
dirs = ["data/mark", "data/yume"]
extract_features_and_train(dirs, mt, mt, st, st, "svm_rbf", "svm_yume_mark")