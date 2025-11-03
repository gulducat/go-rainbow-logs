# 🌈 rainbow 🌈

```
go install github.com/gulducat/go-rainbow-logs/cmd/rainbow

cat <<LOGS | rainbow
    2024-05-13T16:17:17.604-0400 [INFO]  client/gc.go:344: client.gc: marking allocation for GC: alloc_id=77e5e62e-9e61-93a0-bead-0e17de6a1027
    2024-05-13T16:17:17.604-0400 [INFO]  go-hclog@v1.6.2/stdlog.go:60: agent: (runner) received finish
    2024-05-13T16:17:17.839-0400 [DEBUG] client/client.go:2495: client: updated allocations: index=785 total=1 pulled=0 filtered=1
    2024-05-13T16:17:17.839-0400 [DEBUG] client/client.go:2568: client: allocation updates: added=0 removed=0 updated=0 ignored=1
    2024-05-13T16:17:17.839-0400 [DEBUG] client/client.go:2612: client: allocation updates applied: added=0 removed=0 updated=0 ignored=1 errors=0
LOGS
```
