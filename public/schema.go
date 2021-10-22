package public

import "github.com/hashicorp/go-memdb"

// 标签表的Schema
func GetTagsSchema() *memdb.DBSchema {
	tagsSchema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"memberTags": &memdb.TableSchema{
				Name: "memberTags",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name: "id",
						Unique: true,
						Indexer: &memdb.StringFieldIndex{Field: "Consumer_id"},
					},
					//Consumer_id
					"dp_dc": &memdb.IndexSchema{
						Name: "dp_dc",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Dp_dc"},
					},
					"zf_fs": &memdb.IndexSchema{
						Name: "zf_fs",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Zf_fs"},
					},
					"fy_sl": &memdb.IndexSchema{
						Name: "fy_sl",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Fy_sl"},
					},
					"fy_nl": &memdb.IndexSchema{
						Name: "fy_nl",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Fy_nl"},
					},
					"jy_sl": &memdb.IndexSchema{
						Name: "jy_sl",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Jy_sl"},
					},
					"jy_pz": &memdb.IndexSchema{
						Name: "jy_pz",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Jy_pz"},
					},
					"pp_zc": &memdb.IndexSchema{
						Name: "pp_zc",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Pp_zc"},
					},
					"sc_lx": &memdb.IndexSchema{
						Name: "sc_lx",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Sc_lx"},
					},
					"jy_jy": &memdb.IndexSchema{
						Name: "jy_jy",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Jy_jy"},
					},
					"jy_zx": &memdb.IndexSchema{
						Name: "jy_zx",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Jy_zx"},
					},
					"jy_bz": &memdb.IndexSchema{
						Name: "jy_bz",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Jy_bz"},
					},
					"jy_lx": &memdb.IndexSchema{
						Name: "jy_lx",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Jy_lx"},
					},
					"hd_qy": &memdb.IndexSchema{
						Name: "hd_qy",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Hd_qy"},
					},
					"hd_cc": &memdb.IndexSchema{
						Name: "hd_cc",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Hd_cc"},
					},
					"gk_lx": &memdb.IndexSchema{
						Name: "gk_lx",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Gk_lx"},
					},
					"hy_lx": &memdb.IndexSchema{
						Name: "hy_lx",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Hy_lx"},
					},
					"cg_lx": &memdb.IndexSchema{
						Name: "cg_lx",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Cg_lx"},
					},
					"gc_sj": &memdb.IndexSchema{
						Name: "gc_sj",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Gc_sj"},
					},
					"gw_zm": &memdb.IndexSchema{
						Name: "gw_zm",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Gw_zm"},
					},
					"gw_sx": &memdb.IndexSchema{
						Name: "gw_sx",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Gw_sx"},
					},
					"xf_hz": &memdb.IndexSchema{
						Name: "xf_hz",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Xf_hz"},
					},
					"xf_jr": &memdb.IndexSchema{
						Name: "xf_jr",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Xf_jr"},
					},
					"xf_dz": &memdb.IndexSchema{
						Name: "xf_dz",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Xf_dz"},
					},
					"xf_ss": &memdb.IndexSchema{
						Name: "xf_ss",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Xf_ss"},
					},
					"xf_nx": &memdb.IndexSchema{
						Name: "xf_nx",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Xf_nx"},
					},
					"xf_sl": &memdb.IndexSchema{
						Name: "xf_sl",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Xf_sl"},
					},
					"xf_yx": &memdb.IndexSchema{
						Name: "xf_yx",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Xf_yx"},
					},
					"xf_sq": &memdb.IndexSchema{
						Name: "xf_sq",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Xf_sq"},
					},
					"xf_cs": &memdb.IndexSchema{
						Name: "xf_cs",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Xf_cs"},
					},
					"hd_cbd": &memdb.IndexSchema{
						Name: "hd_cbd",
						Unique: false,
						AllowMissing: true,
						Indexer: &memdb.IntFieldIndex{Field: "Hd_cbd"},
					},
				},
			},
		},
	}
	return tagsSchema
}

// 消费者标签结构
type Tags struct {
	Consumer_id   	string `json:"consumer_id"`
	Dp_dc    		int `json:"dp_dc"`
	Zf_fs    		int `json:"zf_fs"`
	Fy_sl    		int `json:"fy_sl"`
	Fy_nl    		int `json:"fy_nl"`
	Jy_sl    		int `json:"jy_sl"`
	Jy_pz    		int `json:"jy_pz"`
	Pp_zc    		int `json:"pp_zc"`
	Sc_lx    		int `json:"sc_lx"`
	Jy_jy    		int `json:"jy_jy"`
	Jy_zx    		int `json:"jy_zx"`
	Jy_bz    		int `json:"jy_bz"`
	Jy_lx    		int `json:"jy_lx"`
	Hd_qy    		int `json:"hd_qy"`
	Hd_cc    		int `json:"hd_cc"`
	Gk_lx    		int `json:"gk_lx"`
	Hy_lx    		int `json:"hy_lx"`
	Cg_lx    		int `json:"cg_lx"`
	Gc_sj    		int `json:"gc_sj"`
	Gw_zm    		int `json:"gw_zm"`
	Gw_sx    		int `json:"gw_sx"`
	Xf_hz    		int `json:"xf_hz"`
	Xf_jr    		int `json:"xf_jr"`
	Xf_dz    		int `json:"xf_dz"`
	Xf_ss    		int `json:"xf_ss"`
	Xf_nx    		int `json:"xf_nx"`
	Xf_sl    		int `json:"xf_sl"`
	Xf_yx    		int `json:"xf_yx"`
	Xf_sq    		int `json:"xf_sq"`
	Xf_cs    		int `json:"xf_cs"`
	Hd_cbd    		int `json:"hd_cbd"`
}