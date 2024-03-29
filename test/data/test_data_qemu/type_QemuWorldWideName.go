package test_data_qemu

func QemuWorldWideName_Legal() []string {
	return []string{
		"",
		"0x8F14641E7F4B27D6",
		"0x43135F8B89D7E9D9",
		"0x0C92E69424B1609D",
		"0x55049C1C8DA23E50",
		"0x9F70557F3A5D09BC",
		"0x2C04D021E89E41FD",
		"0x3D7071A0976895A0",
		"0x7914EAC8244DB5A2",
		"0x30A8A0D6A24B9563",
		"0x1DA1CDD690CE41FE",
		"0x79D8D06D781F8D16",
		"0x821F901FEFEA690F",
		"0x13A9628C95F19A4F",
		"0x6D9E42EFDE9EB1D1",
		"0x2D7167DFCD4C5B48",
		"0x5F792FE5AD87513D",
		"0x9B9B9E1508D0B8B9",
		"0x1C179E26A7B1A040",
		"0x4C5DB30D3E9B1C0A",
		"0x7AED3B1457078D6A",
		"0x6F64A1761AB78825",
		"0x5D8EDD9C5B95C478",
		"0x3C2857A432EC3E9B",
		"0x1EE0843512C6B6BE",
		"0x45272E11136D0B4F",
		"0x08C7D3C6E51EF776",
		"0x28B3B1B41D53A770",
		"0x4A937430E15BCE2B",
		"0x183B7321B0F62A80",
		"0x38C24563EAD1C289",
		"0x5B1DAF347BEEF152",
		"0x0E60D51E5CB50D1A",
		"0x30B2393E28A3EDE4",
		"0x47954C1E3B5CAB46",
		"0x68C0D3D016D68C23",
		"0x933BB8B523C8ACF6",
		"0x1F25C742A9FBE60D",
		"0x09B3EFEF2C88FFAC",
		"0x8B0D9F41A15B7FA6",
		"0x6E2A86B5E950857C",
		"0x3F90A27347DC4F36",
		"0x17619BC0F9F02DBB",
		"0x31D591AE875E8712",
		"0x787E9C1D6D25C12F",
		"0x586AF239D285F6D5",
		"0x961DDAA98CC9A58C",
		"0x241FD8850C9CCDF2",
		"0x9A4D8EAEBC50E3A9",
		"0x6A33FEC7390CFC45",
		"0x1B3CE82A3F2C8EFA",
		"0x42C22809DB71A0E1",
	}
}

func QemuWorldWideName_Illegal() []string {
	return []string{
		"0x44556677EEFF",
		"0x09876:54321",
		"0x1AB2:3C4D",
		"586AF239D285F6D5,",
		"0x0A0B0C0D0E0F",
		"0x12:34:56:78",
		"0x98:76:54:32",
		"0x00:11:22:33:44:55:66:77",
		"0xAA:BB:CC:DD:EE:FF",
		"0x1:2:3:4:5:6:7",
		"0122334455667788",
		"x9A9B9C9D9E9F",
		"0x1A1B1C1D1E1F",
		"0x11223344AABBCC",
		"x44556677",
		"0x0987:6543",
		"0x1A:B2:3C:4D",
		"0x55667788",
		"0x0A:0B:0C:0D",
		"0x12:34:56",
		"0x98:76:54",
		"0x00:11:22:33:44:55:66",
		"0xAA:BB:CC:DD:EE",
		"1:2:3:4:5:6",
		"0x11223344556677",
		"9A9B9C9D9E9",
		"0x1A1B1C1D1E",
		"0x11223344AABBC",
		"0x44556677EE",
		"0x098765:54321",
		"0x1AB:2C4D",
		"0x556677889",
		"0x0A0B0C0D0",
		"12:34:5:78",
		"0x98:76:54:3",
		"0x00:11:22:33:44:55:6",
		"0xAA:BB:CC:DD:EE:F",
		"1:2:3:4:5:6:7",
		"0x1122334455667",
		"0x9A9B9C9D9E",
		"0x1A1B1C1D1",
		"11223344AABB",
		"0x44556677EEF",
		"0x09876:5432",
		"0x1AB2:C4D",
		"0x5566778899,",
		"0x0A0B0C0D",
		"0x12:34:56:7",
		"x98:76:54:2",
	}
}
