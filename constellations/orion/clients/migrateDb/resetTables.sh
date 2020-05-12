#!/bin/bash
migrate -source file://pkg/repos/migrations -database mysql://root:localpass@\(localhost\)/mathnavdb down
migrate -source file://pkg/repos/migrations -database mysql://root:localpass@\(localhost\)/mathnavdb up