const rawDiff = `diff --git a/a/prefix.txt b/a/old_prefix.txt
similarity index 100%
rename from a/prefix.txt
rename to a/old_prefix.txt
diff --git a/b/new_prefix.txt b/b/new_prefix.txt
index 0000000..e9956ba
+++ b/b/new_prefix.txt
+/b prefix
diff --git a/b/prefix.txt b/b/prefix.txt
deleted file mode 100644
index da5f6ce..0000000
--- a/b/prefix.txt
+++ /dev/null
@@ -1 +0,0 @@
-b prefix