package tstorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partitionList_Remove(t *testing.T) {
	tests := []struct {
		name              string
		partitionList     partitionListImpl
		target            partition
		wantErr           bool
		wantPartitionList partitionListImpl
	}{
		{
			name:          "empty partition",
			partitionList: partitionListImpl{},
			wantErr:       true,
		},
		{
			name: "remove the head node",
			partitionList: func() partitionListImpl {
				second := &partitionNode{
					val: &fakePartition{
						minTimestamp: 2,
					},
				}

				first := &partitionNode{
					val: &fakePartition{
						minTimestamp: 1,
					},
					next: second,
				}
				return partitionListImpl{
					size: 2,
					head: first,
					tail: second,
				}
			}(),
			target: &fakePartition{
				minTimestamp: 1,
			},
			wantPartitionList: partitionListImpl{
				size: 1,
				head: &partitionNode{
					val: &fakePartition{
						minTimestamp: 2,
					},
				},
				tail: &partitionNode{
					val: &fakePartition{
						minTimestamp: 2,
					},
				},
			},
		},
		{
			name: "remove the tail node",
			partitionList: func() partitionListImpl {
				second := &partitionNode{
					val: &fakePartition{
						minTimestamp: 2,
					},
				}

				first := &partitionNode{
					val: &fakePartition{
						minTimestamp: 1,
					},
					next: second,
				}
				return partitionListImpl{
					size: 2,
					head: first,
					tail: second,
				}
			}(),
			target: &fakePartition{
				minTimestamp: 2,
			},
			wantPartitionList: partitionListImpl{
				size: 1,
				head: &partitionNode{
					val: &fakePartition{
						minTimestamp: 1,
					},
				},
				tail: &partitionNode{
					val: &fakePartition{
						minTimestamp: 1,
					},
				},
			},
		},
		{
			name: "remove the middle node",
			partitionList: func() partitionListImpl {
				third := &partitionNode{
					val: &fakePartition{
						minTimestamp: 3,
					},
				}
				second := &partitionNode{
					val: &fakePartition{
						minTimestamp: 2,
					},
					next: third,
				}
				first := &partitionNode{
					val: &fakePartition{
						minTimestamp: 1,
					},
					next: second,
				}
				return partitionListImpl{
					size: 3,
					head: first,
					tail: third,
				}
			}(),
			target: &fakePartition{
				minTimestamp: 2,
			},
			wantPartitionList: partitionListImpl{
				size: 2,
				head: &partitionNode{
					val: &fakePartition{
						minTimestamp: 1,
					},
					next: &partitionNode{
						val: &fakePartition{
							minTimestamp: 3,
						},
					},
				},
				tail: &partitionNode{
					val: &fakePartition{
						minTimestamp: 3,
					},
				},
			},
		},
		{
			name: "given node not found",
			partitionList: func() partitionListImpl {
				second := &partitionNode{
					val: &fakePartition{
						minTimestamp: 2,
					},
				}

				first := &partitionNode{
					val: &fakePartition{
						minTimestamp: 1,
					},
					next: second,
				}
				return partitionListImpl{
					size: 2,
					head: first,
					tail: second,
				}
			}(),
			target: &fakePartition{
				minTimestamp: 3,
			},
			wantPartitionList: func() partitionListImpl {
				second := &partitionNode{
					val: &fakePartition{
						minTimestamp: 2,
					},
				}

				first := &partitionNode{
					val: &fakePartition{
						minTimestamp: 1,
					},
					next: second,
				}
				return partitionListImpl{
					size: 2,
					head: first,
					tail: second,
				}
			}(),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.partitionList.Remove(tt.target)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantPartitionList, tt.partitionList)
		})
	}
}

func Test_partitionList_Swap(t *testing.T) {
	tests := []struct {
		name              string
		partitionList     partitionListImpl
		old               partition
		new               partition
		wantErr           bool
		wantPartitionList partitionListImpl
	}{
		{
			name:          "empty partition",
			partitionList: partitionListImpl{},
			wantErr:       true,
		},
		{
			name: "swap the head node",
			partitionList: func() partitionListImpl {
				second := &partitionNode{
					val: &fakePartition{
						minTimestamp: 2,
					},
				}

				first := &partitionNode{
					val: &fakePartition{
						minTimestamp: 1,
					},
					next: second,
				}
				return partitionListImpl{
					size: 2,
					head: first,
					tail: second,
				}
			}(),
			old: &fakePartition{
				minTimestamp: 1,
			},
			new: &fakePartition{
				minTimestamp: 100,
			},
			wantPartitionList: partitionListImpl{
				size: 2,
				head: &partitionNode{
					val: &fakePartition{
						minTimestamp: 100,
					},
					next: &partitionNode{
						val: &fakePartition{
							minTimestamp: 2,
						},
					},
				},
				tail: &partitionNode{
					val: &fakePartition{
						minTimestamp: 2,
					},
				},
			},
		},
		{
			name: "swap the tail node",
			partitionList: func() partitionListImpl {
				second := &partitionNode{
					val: &fakePartition{
						minTimestamp: 2,
					},
				}

				first := &partitionNode{
					val: &fakePartition{
						minTimestamp: 1,
					},
					next: second,
				}
				return partitionListImpl{
					size: 2,
					head: first,
					tail: second,
				}
			}(),
			old: &fakePartition{
				minTimestamp: 2,
			},
			new: &fakePartition{
				minTimestamp: 100,
			},
			wantPartitionList: partitionListImpl{
				size: 2,
				head: &partitionNode{
					val: &fakePartition{
						minTimestamp: 1,
					},
					next: &partitionNode{
						val: &fakePartition{
							minTimestamp: 100,
						},
					},
				},
				tail: &partitionNode{
					val: &fakePartition{
						minTimestamp: 100,
					},
				},
			},
		},
		{
			name: "swap the middle node",
			partitionList: func() partitionListImpl {
				third := &partitionNode{
					val: &fakePartition{
						minTimestamp: 3,
					},
				}
				second := &partitionNode{
					val: &fakePartition{
						minTimestamp: 2,
					},
					next: third,
				}

				first := &partitionNode{
					val: &fakePartition{
						minTimestamp: 1,
					},
					next: second,
				}
				return partitionListImpl{
					size: 3,
					head: first,
					tail: third,
				}
			}(),
			old: &fakePartition{
				minTimestamp: 2,
			},
			new: &fakePartition{
				minTimestamp: 100,
			},
			wantPartitionList: partitionListImpl{
				size: 3,
				head: &partitionNode{
					val: &fakePartition{
						minTimestamp: 1,
					},
					next: &partitionNode{
						val: &fakePartition{
							minTimestamp: 100,
						},
						next: &partitionNode{
							val: &fakePartition{
								minTimestamp: 3,
							},
						},
					},
				},
				tail: &partitionNode{
					val: &fakePartition{
						minTimestamp: 3,
					},
				},
			},
		},
		{
			name: "given node not found",
			partitionList: func() partitionListImpl {
				second := &partitionNode{
					val: &fakePartition{
						minTimestamp: 2,
					},
				}

				first := &partitionNode{
					val: &fakePartition{
						minTimestamp: 1,
					},
					next: second,
				}
				return partitionListImpl{
					size: 2,
					head: first,
					tail: second,
				}
			}(),
			old: &fakePartition{
				minTimestamp: 100,
			},
			wantPartitionList: partitionListImpl{
				size: 2,
				head: &partitionNode{
					val: &fakePartition{
						minTimestamp: 1,
					},
					next: &partitionNode{
						val: &fakePartition{
							minTimestamp: 2,
						},
					},
				},
				tail: &partitionNode{
					val: &fakePartition{
						minTimestamp: 2,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.partitionList.Swap(tt.old, tt.new)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantPartitionList, tt.partitionList)
			/*
				// Check if the partition list is as same as we'd like
				iterator := tt.partitionListImpl.NewIterator()
				partitons := make([]partition, 0, tt.partitionListImpl.Size())
				for iterator.Next() {
					v, err := iterator.Value()
					assert.NoError(t, err)
					partitons = append(partitons, v)
				}
				assert.Equal(t, tt.wantPartitions, partitons)
			*/
		})
	}
}
