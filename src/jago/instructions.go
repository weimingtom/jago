package jago


type InstructionRegistry []Instruction

func (this InstructionRegistry) GetInstruction(opcode uint8) Instruction {
	return this[opcode]
}

func (this InstructionRegistry) RegisterInstructions()  {
	instructions := []Instruction{ // JVM_OPC_MAX+1
		/* ----- CONSTANTS -----------*/
		/*00 (0x00)*/ Instruction{"nop", NOP},
		/*01 (0x01)*/ Instruction{"aconst_null", ACONST_NULL},
		/*02 (0x02)*/ Instruction{"iconst_m1", ICONST_M1},
		/*03 (0x03)*/ Instruction{"iconst_0", ICONST_0},
		/*04 (0x04)*/ Instruction{"iconst_1", ICONST_1},
		/*05 (0x05)*/ Instruction{"iconst_2", ICONST_2},
		/*06 (0x06)*/ Instruction{"iconst_3", ICONST_3},
		/*07 (0x07)*/ Instruction{"iconst_4", ICONST_4},
		/*08 (0x08)*/ Instruction{"iconst_5", ICONST_5},
		/*09 (0x09)*/ Instruction{"lconst_0", LCONST_0},
		/*10 (0x0A)*/ Instruction{"lconst_1", LCONST_1},
		/*11 (0x0B)*/ Instruction{"fconst_0", FCONST_0},
		/*12 (0x0C)*/ Instruction{"fconst_1", FCONST_1},
		/*13 (0x0D)*/ Instruction{"fconst_2", FCONST_2},
		/*14 (0x0E)*/ Instruction{"dconst_0", DCONST_0},
		/*15 (0x0F)*/ Instruction{"dconst_1", DCONST_1},
		/*16 (0x10)*/ Instruction{"bipush", BIPUSH},
		/*17 (0x11)*/ Instruction{"sipush", SIPUSH},
		/*18 (0x12)*/ Instruction{"ldc", LDC},
		/*19 (0x13)*/ Instruction{"ldc_w", LDC_W},
		/*20 (0x14)*/ Instruction{"ldc2_w", LDC2_W},
		/*--------LOADS ----------------*/
		/*21 (0x15)*/ Instruction{"iload", ILOAD},
		/*22 (0x16)*/ Instruction{"lload", LLOAD},
		/*23 (0x17)*/ Instruction{"fload", FLOAD},
		/*24 (0x18)*/ Instruction{"dload", DLOAD},
		/*25 (0x19)*/ Instruction{"aload", ALOAD},
		/*26 (0x1A)*/ Instruction{"iload_0", ILOAD_0},
		/*27 (0x1B)*/ Instruction{"iload_1", ILOAD_1},
		/*28 (0x1C)*/ Instruction{"iload_2", ILOAD_2},
		/*29 (0x1D)*/ Instruction{"iload_3", ILOAD_3},
		/*30 (0x1E)*/ Instruction{"lload_0", LLOAD_0},
		/*31 (0x1F)*/ Instruction{"lload_1", LLOAD_1},
		/*32 (0x20)*/ Instruction{"lload_2", LLOAD_2},
		/*33 (0x21)*/ Instruction{"lload_3", LLOAD_3},
		/*34 (0x22)*/ Instruction{"fload_0", FLOAD_0},
		/*35 (0x23)*/ Instruction{"fload_1", FLOAD_1},
		/*36 (0x24)*/ Instruction{"fload_2", FLOAD_2},
		/*37 (0x25)*/ Instruction{"fload_3", FLOAD_3},
		/*38 (0x26)*/ Instruction{"dload_0", DLOAD_0},
		/*39 (0x27)*/ Instruction{"dload_1", DLOAD_1},
		/*40 (0x28)*/ Instruction{"dload_2", DLOAD_2},
		/*41 (0x29)*/ Instruction{"dload_3", DLOAD_3},
		/*42 (0x2A)*/ Instruction{"aload_0", ALOAD_0},
		/*43 (0x2B)*/ Instruction{"aload_1", ALOAD_1},
		/*44 (0x2C)*/ Instruction{"aload_2", ALOAD_2},
		/*45 (0x2D)*/ Instruction{"aload_3", ALOAD_3},
		/*46 (0x2E)*/ Instruction{"iaload", IALOAD},
		/*47 (0x2F)*/ Instruction{"laload", LALOAD},
		/*48 (0x30)*/ Instruction{"faload", FALOAD},
		/*49 (0x31)*/ Instruction{"daload", DALOAD},
		/*50 (0x32)*/ Instruction{"aaload", AALOAD},
		/*51 (0x33)*/ Instruction{"baload", BALOAD},
		/*52 (0x34)*/ Instruction{"caload", CALOAD},
		/*53 (0x35)*/ Instruction{"saload", SALOAD},
		/*--------STORES ----------------*/
		/*54 (0x36)*/ Instruction{"istore", ISTORE},
		/*55 (0x37)*/ Instruction{"lstore", LSTORE},
		/*56 (0x38)*/ Instruction{"fstore", FSTORE},
		/*57 (0x39)*/ Instruction{"dstore", DSTORE},
		/*58 (0x3A)*/ Instruction{"astore", ASTORE},
		/*59 (0x3B)*/ Instruction{"istore_0", ISTORE_0},
		/*60 (0x3C)*/ Instruction{"istore_1", ISTORE_1},
		/*61 (0x3D)*/ Instruction{"istore_2", ISTORE_2},
		/*62 (0x3E)*/ Instruction{"istore_3", ISTORE_3},
		/*63 (0x3F)*/ Instruction{"lstore_0", LSTORE_0},
		/*64 (0x40)*/ Instruction{"lstore_1", LSTORE_1},
		/*65 (0x41)*/ Instruction{"lstore_2", LSTORE_2},
		/*66 (0x42)*/ Instruction{"lstore_3", LSTORE_3},
		/*67 (0x43)*/ Instruction{"fstore_0", FSTORE_0},
		/*68 (0x44)*/ Instruction{"fstore_1", FSTORE_1},
		/*69 (0x45)*/ Instruction{"fstore_2", FSTORE_2},
		/*70 (0x46)*/ Instruction{"fstore_3", FSTORE_3},
		/*71 (0x47)*/ Instruction{"dstore_0", DSTORE_0},
		/*72 (0x48)*/ Instruction{"dstore_1", DSTORE_1},
		/*73 (0x49)*/ Instruction{"dstore_2", DSTORE_2},
		/*74 (0x4A)*/ Instruction{"dstore_3", DSTORE_3},
		/*75 (0x4B)*/ Instruction{"astore_0", ASTORE_0},
		/*76 (0x4C)*/ Instruction{"astore_1", ASTORE_1},
		/*77 (0x4D)*/ Instruction{"astore_2", ASTORE_2},
		/*78 (0x4E)*/ Instruction{"astore_3", ASTORE_3},
		/*79 (0x4F)*/ Instruction{"iastore", IASTORE},
		/*80 (0x50)*/ Instruction{"lastore", LASTORE},
		/*81 (0x51)*/ Instruction{"fastore", FASTORE},
		/*82 (0x52)*/ Instruction{"dastore", DASTORE},
		/*83 (0x53)*/ Instruction{"aastore", AASTORE},
		/*84 (0x54)*/ Instruction{"bastore", BASTORE},
		/*85 (0x55)*/ Instruction{"castore", CASTORE},
		/*86 (0x56)*/ Instruction{"sastore", SASTORE},
		/*--------STACK---------------*/
		/*87 (0x57)*/ Instruction{"pop", POP},
		/*88 (0x58)*/ Instruction{"pop2", POP2},
		/*89 (0x59)*/ Instruction{"dup", DUP},
		/*90 (0x5A)*/ Instruction{"dup_x1", DUP_X1},
		/*91 (0x5B)*/ Instruction{"dup_x2", DUP_X2},
		/*92 (0x5C)*/ Instruction{"dup2", DUP2},
		/*93 (0x5D)*/ Instruction{"dup2_x1", DUP2_X1},
		/*94 (0x5E)*/ Instruction{"dup2_x2", DUP2_X2},
		/*95 (0x5F)*/ Instruction{"swap", SWAP},
		/*---------MATH -------------*/
		/*96 (0x60)*/  Instruction{"iadd", IADD},
		/*97 (0x61)*/  Instruction{"ladd", LADD},
		/*98 (0x62)*/  Instruction{"fadd", FADD},
		/*99 (0x63)*/  Instruction{"dadd", DADD},
		/*100 (0x64)*/ Instruction{"isub", ISUB},
		/*101 (0x65)*/ Instruction{"lsub", LSUB},
		/*102 (0x66)*/ Instruction{"fsub", FSUB},
		/*103 (0x67)*/ Instruction{"dsub", DSUB},
		/*104 (0x68)*/ Instruction{"imul", IMUL},
		/*105 (0x69)*/ Instruction{"lmul", LMUL},
		/*106 (0x6A)*/ Instruction{"fmul", FMUL},
		/*107 (0x6B)*/ Instruction{"dmul", DMUL},
		/*108 (0x6C)*/ Instruction{"idiv", IDIV},
		/*109 (0x6D)*/ Instruction{"ldiv", LDIV},
		/*110 (0x6E)*/ Instruction{"fdiv", FDIV},
		/*111 (0x6F)*/ Instruction{"ddiv", DDIV},
		/*112 (0x70)*/ Instruction{"irem", IREM},
		/*113 (0x71)*/ Instruction{"lrem", LREM},
		/*114 (0x72)*/ Instruction{"frem", FREM},
		/*115 (0x73)*/ Instruction{"drem", DREM},
		/*116 (0x74)*/ Instruction{"ineg", INEG},
		/*117 (0x75)*/ Instruction{"lneg", LNEG},
		/*118 (0x76)*/ Instruction{"fneg", FNEG},
		/*119 (0x77)*/ Instruction{"dneg", DNEG},
		/*120 (0x78)*/ Instruction{"ishl", ISHL},
		/*121 (0x79)*/ Instruction{"lshl", LSHL},
		/*122 (0x7A)*/ Instruction{"ishr", ISHR},
		/*123 (0x7B)*/ Instruction{"lshr", LSHR},
		/*124 (0x7C)*/ Instruction{"iushr", IUSHR},
		/*125 (0x7D)*/ Instruction{"lushr", LUSHR},
		/*126 (0x7E)*/ Instruction{"iand", IAND},
		/*127 (0x7F)*/ Instruction{"land", LAND},
		/*128 (0x80)*/ Instruction{"ior", IOR},
		/*129 (0x81)*/ Instruction{"lor", LOR},
		/*130 (0x82)*/ Instruction{"ixor", IXOR},
		/*131 (0x83)*/ Instruction{"lxor", LXOR},
		/*132 (0x84)*/ Instruction{"iinc", IINC},
		/*--------CONVERSIONS-----------*/
		/*133 (0x85)*/ Instruction{"i2l", I2L},
		/*134 (0x86)*/ Instruction{"i2f", I2F},
		/*135 (0x87)*/ Instruction{"i2d", I2D},
		/*136 (0x88)*/ Instruction{"l2i", L2I},
		/*137 (0x89)*/ Instruction{"l2f", L2F},
		/*138 (0x8A)*/ Instruction{"l2d", L2D},
		/*139 (0x8B)*/ Instruction{"f2i", F2I},
		/*140 (0x8C)*/ Instruction{"f2l", F2L},
		/*141 (0x8D)*/ Instruction{"f2d", F2D},
		/*142 (0x8E)*/ Instruction{"d2i", D2I},
		/*143 (0x8F)*/ Instruction{"d2l", D2L},
		/*144 (0x90)*/ Instruction{"d2f", D2F},
		/*145 (0x91)*/ Instruction{"i2b", I2B},
		/*146 (0x92)*/ Instruction{"i2c", I2C},
		/*147 (0x93)*/ Instruction{"i2s", I2S},
		/*-----------COMPARASON -----------*/
		/*148 (0x94)*/ Instruction{"lcmp", LCMP},
		/*149 (0x95)*/ Instruction{"fcmpl", FCMPL},
		/*150 (0x96)*/ Instruction{"fcmpg", FCMPG},
		/*151 (0x97)*/ Instruction{"dcmpl", DCMPL},
		/*152 (0x98)*/ Instruction{"dcmpg", DCMPG},
		/*153 (0x99)*/ Instruction{"ifeq", IFEQ},
		/*154 (0x9A)*/ Instruction{"ifne", IFNE},
		/*155 (0x9B)*/ Instruction{"iflt", IFLT},
		/*156 (0x9C)*/ Instruction{"ifge", IFGE},
		/*157 (0x9D)*/ Instruction{"ifgt", IFGT},
		/*158 (0x9E)*/ Instruction{"ifle", IFLE},
		/*159 (0x9F)*/ Instruction{"if_icmpeq", IF_ICMPEQ},
		/*160 (0xA0)*/ Instruction{"if_icmpne", IF_ICMPNE},
		/*161 (0xA1)*/ Instruction{"if_icmplt", IF_ICMPLT},
		/*162 (0xA2)*/ Instruction{"if_icmpge", IF_ICMPGE},
		/*163 (0xA3)*/ Instruction{"if_icmpgt", IF_ICMPGT},
		/*164 (0xA4)*/ Instruction{"if_icmple", IF_ICMPLE},
		/*165 (0xA5)*/ Instruction{"if_acmpeq", IF_ACMPEQ},
		/*166 (0xA6)*/ Instruction{"if_acmpne", IF_ACMPNE},
		/*---------REFERENCES -------------*/
		/*167 (0xA7)*/ Instruction{"goto", GOTO},
		/*168 (0xA8)*/ Instruction{"jsr", JSR},
		/*169 (0xA9)*/ Instruction{"ret", RET},
		/*170 (0xAA)*/ Instruction{"tableswitch", TABLESWITCH}, // variable bytecode length
		/*171 (0xAB)*/ Instruction{"lookupswitch", LOOKUPSWITCH},
		/*172 (0xAC)*/ Instruction{"ireturn", IRETURN},
		/*173 (0xAD)*/ Instruction{"lreturn", LRETURN},
		/*174 (0xAE)*/ Instruction{"freturn", FRETURN},
		/*175 (0xAF)*/ Instruction{"dreturn", DRETURN},
		/*176 (0xB0)*/ Instruction{"areturn", ARETURN},
		/*177 (0xB1)*/ Instruction{"return", RETURN},
		/*-------CONTROLS------------------*/
		/*178 (0xB2)*/ Instruction{"getstatic", GETSTATIC},
		/*179 (0xB3)*/ Instruction{"putstatic", PUTSTATIC},
		/*180 (0xB4)*/ Instruction{"getfield", GETFIELD},
		/*181 (0xB5)*/ Instruction{"putfield", PUTFIELD},
		/*182 (0xB6)*/ Instruction{"invokevirtual", INVOKEVIRTUAL},
		/*183 (0xB7)*/ Instruction{"invokespecial", INVOKESPECIAL},
		/*184 (0xB8)*/ Instruction{"invokestatic", INVOKESTATIC},
		/*185 (0xB9)*/ Instruction{"invokeinterface", INVOKEINTERFACE},
		/*186 (0xBA)*/ Instruction{"invokedynamic", INVOKEDYNAMIC},
		/*187 (0xBB)*/ Instruction{"new", NEW},
		/*188 (0xBC)*/ Instruction{"newarray", NEWARRAY},
		/*189 (0xBD)*/ Instruction{"anewarray", ANEWARRAY},
		/*190 (0xBE)*/ Instruction{"arraylength", ARRAYLENGTH},
		/*191 (0xBF)*/ Instruction{"athrow", ATHROW},
		/*192 (0xC0)*/ Instruction{"checkcast", CHECKCAST},
		/*193 (0xC1)*/ Instruction{"instanceof", INSTANCEOF},
		/*194 (0xC2)*/ Instruction{"monitorenter", MONITORENTER},
		/*195 (0xC3)*/ Instruction{"monitorexit", MONITOREXIT},
		/*--------EXTENDED-----------------*/
		/*196 (0xC4)*/ Instruction{"wide", WIDE},
		/*197 (0xC5)*/ Instruction{"multianewarray", MULTIANEWARRAY},
		/*198 (0xC6)*/ Instruction{"ifnull", IFNULL},
		/*199 (0xC7)*/ Instruction{"ifnonnull", IFNONNULL},
		/*200 (0xC8)*/ Instruction{"goto_w", GOTO_W},
		/*201 (0xC9)*/ Instruction{"jsr_w", JSR_W},
		/*----------RESERVED ---------------*/
		/*202 (0xCA)*/ Instruction{"breakpoint", BREAKPOINT},
		///*254 (0xFE)*/    Instruction{"impdep1", IMPDEP1},
		///*255 (0xFF)*/    Instruction{"impdep2", IMPDEP2},
	}

	for i, instruction := range instructions {
		this[i] = instruction
	}
}

type Instruction struct {
	mnemonic  string
	interpret func(*Thread, *Frame, *Class, *Method)
}

func assertIntCompatible(value Value) Value {
	switch value.(type) {
	case Boolean, Byte, Short, Char, Int:
	default:
		Fatal("%s is not int compatible", value.Type().Name())
	}
	return value
}

func intCompatible(value Value) Int {
	assertIntCompatible(value)

	switch value.(type) {
	case Boolean:
		return Int(value.(Boolean))
	case Byte:
		return Int(value.(Byte))
	case Short:
		return Int(value.(Short))
	case Char:
		return Int(value.(Char))
	case Int:
		return value.(Int)
	default:
		Fatal("%s is not int compatible", value.Type().Name())
	}

	// never be here
	return 0
}


